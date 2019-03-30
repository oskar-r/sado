package middleware

import (
	"errors"
	"log"
	"net/http"
	"reflect"
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

/*
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
		uid := ""
		if v.FieldByName("UserID").IsValid() {
			uid = v.FieldByName("UserID").String()
		}
		if uid == "" {
			resp := map[string]string{"error": "no user id in claim"}
			c.JSON(http.StatusBadRequest, resp)
			c.Abort()
			return
		}

		//Get resource from path
		var ids []int

		//Filter out all policies that the grants access to the user for the resource
		if id == "all" {
			log.Printf("[TEST] %s %s %s %s", uid, obj, id, act)
			ok = e.GetFilteredPolicy(uid, obj, id, act)
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
			ok = e.Enforce(uid, obj, id, act)
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

		c.Set("resources", ids)
		c.Set("userID", uid)
		c.Next()
	}
}
*/

func extractClaim(c *gin.Context) (string, error) {
	claim, ok := c.Get("user")
	if !ok {
		return "", errors.New("missing claim")
	}

	v := reflect.Indirect(reflect.ValueOf(claim))
	uid := ""
	if v.FieldByName("UserID").IsValid() {
		uid = v.FieldByName("UserID").String()
	}
	if uid == "" {
		return "", errors.New("no user id in claim")
	}

	return uid, nil
}
func ActionAuthorizer(e *casbin.Enforcer, id string, obj string, act string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Get user from JWT
		uid, err := extractClaim(c)
		if err != nil {
			resp := map[string]string{"error": err.Error()}
			c.JSON(http.StatusForbidden, resp)
			c.Abort()
			return
		}

		if ok := e.Enforce(uid, obj, "*", act); !ok {
			resp := map[string]string{"error": "no access to resource"}
			c.JSON(http.StatusForbidden, resp)
			c.Abort()
			return
		}
		c.Set("userID", uid)
		c.Next()
	}
}

/*
func ResourceAuthorizer(e *casbin.Enforcer, id string, obj string, act string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Get user from JWT
		uid, err := extractClaim(c)
		if err != nil {
			resp := map[string]string{"error": err.Error()}
			c.JSON(http.StatusForbidden, resp)
			c.Abort()
			return
		}

		//Get resource from path
		var res bool
		var ids []int

		//Filter out all policies that the grants access to the user for the resource

		if id == "all" {
			res, err = e.EnforceSafe(uid, obj, act)
			if err == nil && res {
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
			res, err = e.EnforceSafe(strconv.Itoa(uid), obj, id, act)
			if err == nil && res {
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
		if err != nil {
			resp := map[string]string{"error": err.Error()}
			c.JSON(http.StatusForbidden, resp)
			c.Abort()
			return
		}
		if !res {
			resp := map[string]string{"error": "no access to resource"}
			c.JSON(http.StatusForbidden, resp)
			c.Abort()
			return
		}

		c.Set("handlerInfo", map[string][]int{
			"authorizedResources": ids,
		})
		c.Next()
	}
}

func GetResourceIDS(sub string, obj string, oid string, re *regexp.Regexp, e *casbin.Enforcer) map[int]string {
	var filters []string
	if oid == "*" { //All resources ids that the use have access to
		filters = append(filters, sub, obj)
	} else {
		filters = append(filters, sub, obj, oid)
	}

	i := make(map[int]string, 0)
	for _, v := range e.GetFilteredPolicy(0, filters...) {
		if len(v) == 4 {
			if re.MatchString(v[3]) && v[2] != "*" {
				it, err := strconv.Atoi(v[2])
				if err == nil {
					i[it] = v[3]
				}
			}
		}
	}
	return i
}
*/
