
#!/bin/bash

# Immediately exits if any error occurs during the script
# execution. If not set, an error could occur and the
# script would continue its execution.
set -o errexit

# Main execution:
# - verifies if all environment variables are set
# - runs the SQL code to create user and database
main() {
  init_user_and_db
}

# Performs the initialization in the already-started PostgreSQL
# using the preconfigured POSTGRE_USER user.
init_user_and_db() {
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$ARCH_DB_DATABASE" <<-EOSQL
     GRANT ALL ON SCHEMA $ARCH_SCHEMA TO $DB_USER;
     GRANT ALL ON ALL TABLES IN SCHEMA $ARCH_SCHEMA TO $DB_USER;
EOSQL
 psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POL_DB_DATABASE" <<-EOSQL
     GRANT ALL ON SCHEMA $POL_SCHEMA TO $DB_USER;
     GRANT ALL ON ALL TABLES IN SCHEMA $POL_SCHEMA TO $DB_USER;
EOSQL
}
# Executes the main routine with environment variables
# passed through the command line. We don't use them in
# this script but now you know 🤓
main "$@"