package middleware

import (
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func PathAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {

	return func(c *gin.Context) {
		//claims := jwt.ExtractClaims(c)
		claim, ok := c.Get(identityKey)
		if !ok {
			resp := map[string]string{"error": "No username i token"}
			c.JSON(http.StatusBadRequest, resp)
			c.Abort()
			return
		}
		var role string
		v := reflect.Indirect(reflect.ValueOf(claim))
		if v.FieldByName("CurrentRole").IsValid() && v.FieldByName("CurrentRole").String() != "" {
			role = v.FieldByName("CurrentRole").String()
		}

		// Validate that user still exists
		// To be implemented

		// casbin rule enforcing
		ts := time.Now()
		log.Printf("[INFO] Enforcing policy %s %s %s %s", role, c.Request.URL.Path, "*", c.Request.Method)
		ok = e.Enforce(role, c.Request.URL.Path, "*", c.Request.Method)
		execTime := (float64(time.Now().UnixNano()) - float64(ts.UnixNano())) / 1000000 //Execution time in miliseconds
		log.Printf("[INFO] Enforce call took %f ms", execTime)
		//res, err := e.EnforceSafe(role, c.Request.URL.Path, "*", c.Request.Method)
		if !ok {
			log.Printf("[INFO] Policy deciled")
			resp := map[string]string{"error": "Policy decline"}
			c.JSON(http.StatusForbidden, resp)
			c.Abort()
			return
		}
		c.Next()
	}
}

func ResourceAuthorizer(e *casbin.Enforcer, id string, obj string, act string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Get user from JWT
		claim, ok := c.Get("user")
		if !ok {
			resp := map[string]string{"error": "missing claim"}
			c.JSON(http.StatusBadRequest, resp)
			c.Abort()
			return
		}

		v := reflect.Indirect(reflect.ValueOf(claim))
		uid := 0
		if v.FieldByName("UserID").IsValid() {
			uid = int(v.FieldByName("UserID").Int())
		}
		if uid == 0 {
			resp := map[string]string{"error": "no user id in claim"}
			c.JSON(http.StatusBadRequest, resp)
			c.Abort()
			return
		}

		//Get resource from path
		var ids []int

		//Filter out all policies that the grants access to the user for the resource
		if id == "all" {
			ok = e.Enforce(strconv.Itoa(uid), obj, id, act)
			if ok {
				idi, err := strconv.Atoi(id)
				if err != nil {
					resp := map[string]string{"error": "resource id should be number"}
					c.JSON(http.StatusForbidden, resp)
					c.Abort()
					return
				}
				ids = append(ids, idi)
			}
		} else {
			id := c.Param(id)
			if id == "" {
				resp := map[string]string{"error": "no resource id provided"}
				c.JSON(http.StatusBadRequest, resp)
				c.Abort()
				return
			}
			ok = e.Enforce(strconv.Itoa(uid), obj, id, act)
			if ok {
				idi, err := strconv.Atoi(id)
				if err != nil {
					resp := map[string]string{"error": "resource id should be number"}
					c.JSON(http.StatusForbidden, resp)
					c.Abort()
					return
				}
				ids = append(ids, idi)
			}
		}
		if !ok {
			resp := map[string]string{"error": "no access to resource"}
			c.JSON(http.StatusForbidden, resp)
			c.Abort()
			return
		}

		c.Set("handlerInfo", map[string][]int{
			"id":     ids,
			"userID": []int{uid},
		})
		c.Next()
	}
}

/*
func getResources(c *gin.Context, e *casbin.Enforcer, ar authReq) ([]int, int, error) {
	sub, err := getAttributeFromClaim(ar.claimStore, ar.claimID, c)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	e.GetFilteredPolicy
	ids, err := pc.Get(context.Background(), &ps.Filter{
		Filters: []string{sub, ar.obj, ar.rid, ar.act},
	})

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("could not fetch resources")
	}
	var idList []int
	for id := range ids.Resources {
		idList = append(idList, int(id))
	}
	return idList, 200, nil
}

type authReq struct {
	claimStore string
	claimID    string
	obj        string
	rid        string
	act        string
	ids        []int
	err        error
	returnCode int
}

func FilterOutResources(e *casbin.Enforcer, obj string, rid string, act string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authRequests := []authReq{}
		authRequests = append(authRequests, authReq{"user", "UserID", obj, rid, act, []int{}, nil, 200})
		authRequests = append(authRequests, authReq{"user", "Role", obj, rid, act, []int{}, nil, 200})

		var wg sync.WaitGroup
		wg.Add(len(authRequests))
		for i := range authRequests {
			go func(i int) {
				defer wg.Done()
				ids, rc, err := getResources(c, pc, authRequests[i])
				if err != nil {
					authRequests[i].err = err
					authRequests[i].returnCode = rc
				} else {
					authRequests[i].ids = ids
				}
				return
			}(i)
		}
		wg.Wait()

		for _, v := range authRequests {
			if v.err != nil {
				c.JSON(v.returnCode, map[string]string{"error": v.err.Error()})
				c.Abort()
				return
			}
		}

		var idList []int
		for _, ar := range authRequests {
			idList = append(idList, ar.ids...)
		}
		c.Set("ids", idList)
		c.Next()
	}
}

func getAttributeFromClaim(claimStoreID string, claimID string, c *gin.Context) (string, error) {
	claim, ok := c.Get(claimStoreID)
	if !ok {
		return "", errors.New("missing claim")
	}

	v := reflect.Indirect(reflect.ValueOf(claim))
	var uid string
	if v.FieldByName(claimID).IsValid() {
		switch v.FieldByName(claimID).Interface().(type) {
		case string:
			uid = v.FieldByName(claimID).String()
		case int:
			uid = strconv.Itoa(int(v.FieldByName(claimID).Int()))
		}
	}
	if uid == "" {
		return "", errors.New("no " + claimID + " in claim")
	}
	log.Printf("[INFO] Got %s", uid)
	return uid, nil

}

func returnError(c *gin.Context, err error, statusCode int) {
	c.JSON(statusCode, map[string]string{"error": err.Error()})
	c.Abort()
}
func FilterByRole(e *casbin.Enforcer, obj string, rid string, act string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, ok := c.Get("user")
		if !ok {
			resp := map[string]string{"error": "missing claim"}
			c.JSON(http.StatusBadRequest, resp)
			c.Abort()
			return
		}

		v := reflect.Indirect(reflect.ValueOf(claim))
		var role string
		if v.FieldByName("Role").IsValid() {
			role = v.FieldByName("UserID").String()
		}
		if role == "" {
			resp := map[string]string{"error": "no role in claim"}
			c.JSON(http.StatusBadRequest, resp)
			c.Abort()
			return
		}

		var idList []int
		ids, err := pc.Get(context.Background(), &ps.Filter{
			Filters: []string{role, obj, rid, act},
		})
		if err != nil {
			resp := map[string]string{"error": "could not fetch resources"}
			c.JSON(http.StatusInternalServerError, resp)
			c.Abort()
			return
		}
		for id := range ids.Resources {
			idList = append(idList, int(id))
		}
		c.Set("ids", idList)
		c.Next()
	}
}
*/
