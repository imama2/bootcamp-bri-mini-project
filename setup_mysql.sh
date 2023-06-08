#!/usr/bin/env bash

# Variables to set.
# You can set remove these lines (or set them to empty strings) and you will then be prompted to enter new values during the process.
# **** NOTE: The DB admin credentials are not the same as the DB user credentials provided in the .env file (unless defined the same). ****
ENV_FILE=".env"
DEFAULT_MYSQL_HOST="localhost"
DEFAULT_MYSQL_ADMIN_USERNAME="root"
DEFAULT_MYSQL_ADMIN_PASSWORD="root"

######## DO NOT EDIT BELOW THIS LINE (unless you know what you're doing :) ##########
MAX_MYSQL_DATABASE_NAME_LENGTH=64
MAX_MYSQL_USERNAME_LENGTH=16

declare -A REQUIRED_ENV_VARS=( [DB_HOST]="" [DB_DATABASE]="" [DB_USERNAME]="" [DB_PASSWORD]="" )

# Exit script on first error
set -e

# Pretty colors!
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

echo "Prepeating database..."

# Verify .env file exists
confirm_env_file_exists() {
    while [ ! -f "${ENV_FILE}" ]; do
        if [ -z "${ENV_FILE}" ]; then
            printf "${RED}Error:${NC} Environment file not specified.\n"
        else
            printf "${RED}Error:${NC} Environment file \"$ENV_FILE\" does not exist.\n"
        fi
        echo -n "Enter the environment file name (e.g. \".env\"): "
        read ENV_FILE
        echo ""
    done
}

if [[ -z "$ENV_FILE" || ! -f "${ENV_FILE}" ]]; then
    confirm_env_file_exists
fi

# Verify MySQL admin credentials.
# The user will be re-prompted to enter correct details if the default credentials fail, or if any subsequent attempt fails.
MYSQL_HOST="${DEFAULT_MYSQL_HOST}"
MYSQL_ADMIN_USERNAME="${DEFAULT_MYSQL_ADMIN_USERNAME}"
MYSQL_ADMIN_PASSWORD="${DEFAULT_MYSQL_ADMIN_PASSWORD}"

get_mysql_credentials() {
    MYSQL_HOST="${DEFAULT_MYSQL_HOST}"
    MYSQL_ADMIN_USERNAME="${DEFAULT_MYSQL_ADMIN_USERNAME}"
    MYSQL_ADMIN_PASSWORD="${DEFAULT_MYSQL_ADMIN_PASSWORD}"

    # Host name
    if [ -z "${DEFAULT_MYSQL_HOST}" ]; then
        while [ -z "${MYSQL_HOST}" ]; do
            echo -n "Enter the MySQL host name (e.g. localhost): "
            read MYSQL_HOST
        done
    else
        echo -n "Enter the MySQL host name (press [ENTER] for \"${DEFAULT_MYSQL_HOST}\"): "
        read MYSQL_HOST
        if [ -z "${MYSQL_HOST}" ]; then
            MYSQL_HOST="$DEFAULT_MYSQL_HOST"
        fi
    fi

    # Admin username
    if [ -z "${DEFAULT_MYSQL_ADMIN_USERNAME}" ]; then
        while [ -z "${MYSQL_ADMIN_USERNAME}" ]; do
            echo -n "Enter the MySQL admin username (e.g. root): "
            read MYSQL_ADMIN_USERNAME
        done
    else
        echo -n "Enter the MySQL admin username (press [ENTER] for \"${DEFAULT_MYSQL_ADMIN_USERNAME}\"): "
        read MYSQL_ADMIN_USERNAME
        if [ -z "${MYSQL_ADMIN_USERNAME}" ]; then
            MYSQL_ADMIN_USERNAME="$DEFAULT_MYSQL_ADMIN_USERNAME"
        fi
    fi

    # Admin password
    if [ -z "${DEFAULT_MYSQL_ADMIN_PASSWORD}" ]; then
        while [ -z "${MYSQL_ADMIN_PASSWORD}" ]; do
            echo -n "Enter the MySQL admin password (e.g. secret): "
            read MYSQL_ADMIN_PASSWORD
        done
    else
        echo -n "Enter the MySQL admin password (press [ENTER] for \"${DEFAULT_MYSQL_ADMIN_PASSWORD}\"): "
        read MYSQL_ADMIN_PASSWORD
        if [ -z "${MYSQL_ADMIN_PASSWORD}" ]; then
            MYSQL_ADMIN_PASSWORD="$DEFAULT_MYSQL_ADMIN_PASSWORD"
        fi
    fi
}

if [[ -z "${MYSQL_HOST}" || -z "${MYSQL_ADMIN_USERNAME}" || -z "${MYSQL_ADMIN_PASSWORD}" ]]; then
    get_mysql_credentials
fi

while ! mysql -h "${MYSQL_HOST}" -u "${MYSQL_ADMIN_USERNAME}" -p"${MYSQL_ADMIN_PASSWORD}" -e ";" >/dev/null 2>&1; do
    echo ""
    echo "*** Your default MySQL admin credentials failed verification (or unable to connect to the default host name). ***"
    echo ""

    get_mysql_credentials
done

ERRORS_EXIST=false

# Verify that the required entries in the .env file exist (and are not set to empty strings)
for i in "${!REQUIRED_ENV_VARS[@]}"
do
    :
    if ! grep -qoEx "${i}=.+" "${ENV_FILE}"; then
        printf "${RED}Error:${NC} Missing required line from your .env file: \"$i={your_value}\"\n"
        ERRORS_EXIST=true
    else
        REQUIRED_ENV_VARS[$i]=`grep -oEx "${i}=.+" "${ENV_FILE}" | sed -n -e "s/^.*${i}=//p"`
    fi
done

# Validate env variable values
for i in "${!REQUIRED_ENV_VARS[@]}"
do
    :
    if [ ${i} == "DB_DATABASE" ]; then
        #echo "Checking DB_DATABASE length..."
        if [ ${#REQUIRED_ENV_VARS[$i]} -gt ${MAX_MYSQL_DATABASE_NAME_LENGTH} ]; then
            printf "${RED}Error:${NC} The length of DB_DATABASE set in file \"${ENV_FILE}\" must not exceed ${MAX_MYSQL_DATABASE_NAME_LENGTH} characters in length.\n"
            ERRORS_EXIST=true
        fi
    elif [ ${i} == "DB_USERNAME" ]; then
        #echo "Checking DB_USERNAME length..."
        if [ ${#REQUIRED_ENV_VARS[$i]} -gt ${MAX_MYSQL_USERNAME_LENGTH} ]; then
            printf "${RED}Error:${NC} The length of DB_USERNAME set in file \"${ENV_FILE}\" must not exceed ${MAX_MYSQL_USERNAME_LENGTH} characters in length.\n"
            ERRORS_EXIST=true
        fi
    fi
done

if [ ${ERRORS_EXIST} == true ]; then
    exit 0;
fi

# Create the database
#todo: if database already exists, prompt for delete & continue
mysql -h "${MYSQL_HOST}" -u "${MYSQL_ADMIN_USERNAME}" -p${MYSQL_ADMIN_PASSWORD} -e "DROP DATABASE IF EXISTS \`${REQUIRED_ENV_VARS[DB_DATABASE]}\`;" >/dev/null 2>&1;
mysql -h "${MYSQL_HOST}" -u "${MYSQL_ADMIN_USERNAME}" -p${MYSQL_ADMIN_PASSWORD} -e "CREATE DATABASE IF NOT EXISTS \`${REQUIRED_ENV_VARS[DB_DATABASE]}\`;" >/dev/null 2>&1
printf "Created MySQL database ${GREEN}${REQUIRED_ENV_VARS[DB_DATABASE]}${NC}.\n"

# Create the MySQL user
mysql -h "${MYSQL_HOST}" -u "${MYSQL_ADMIN_USERNAME}" -p${MYSQL_ADMIN_PASSWORD} -e "GRANT USAGE ON *.* TO '${REQUIRED_ENV_VARS[DB_USERNAME]}'@'${REQUIRED_ENV_VARS[DB_HOST]}' IDENTIFIED BY '${REQUIRED_ENV_VARS[DB_PASSWORD]}';" >/dev/null 2>&1;
mysql -h "${MYSQL_HOST}" -u "${MYSQL_ADMIN_USERNAME}" -p${MYSQL_ADMIN_PASSWORD} -e "DROP USER '${REQUIRED_ENV_VARS[DB_USERNAME]}'@'${REQUIRED_ENV_VARS[DB_HOST]}';" >/dev/null 2>&1;
mysql -h "${MYSQL_HOST}" -u "${MYSQL_ADMIN_USERNAME}" -p${MYSQL_ADMIN_PASSWORD} -e "GRANT ALL PRIVILEGES ON ${REQUIRED_ENV_VARS[DB_DATABASE]}.* TO '${REQUIRED_ENV_VARS[DB_USERNAME]}'@'${REQUIRED_ENV_VARS[DB_HOST]}' IDENTIFIED BY '${REQUIRED_ENV_VARS[DB_PASSWORD]}';" >/dev/null 2>&1
mysql -h "${MYSQL_HOST}" -u "${MYSQL_ADMIN_USERNAME}" -p${MYSQL_ADMIN_PASSWORD} -e "FLUSH PRIVILEGES;" >/dev/null 2>&1
printf "Created MySQL user ${GREEN}${REQUIRED_ENV_VARS[DB_USERNAME]}${NC} with password ${GREEN}${REQUIRED_ENV_VARS[DB_PASSWORD]}${NC}.\n"

# Verify MySQL database was created (mysql exit code 0)
mysql -h "${MYSQL_HOST}" -u "${MYSQL_ADMIN_USERNAME}" -p${MYSQL_ADMIN_PASSWORD} -e "use \`${REQUIRED_ENV_VARS[DB_DATABASE]}\`;" >/dev/null 2>&1
if [ $? == "0" ]; then
    printf "${GREEN}Database verified!${NC} -- Verified database ${GREEN}${REQUIRED_ENV_VARS[DB_DATABASE]}${NC} exists on host ${GREEN}${REQUIRED_ENV_VARS[DB_HOST]}${NC}\n"
else
    printf "${RED}Warning: Database could not be verified!${NC} -- Unable to verify that database ${RED}${REQUIRED_ENV_VARS[DB_DATABASE]}${NC} was created on host ${RED}${REQUIRED_ENV_VARS[DB_HOST]}${NC}\n"
fi

# Verify MySQL user was created
MYSQL_USER_EXISTS=false
{
    while read User; do
        if [[ "${REQUIRED_ENV_VARS[DB_USERNAME]}" == "$User" ]]; then
            MYSQL_USER_EXISTS=true
            break
        fi
    done < <(mysql -h "${MYSQL_HOST}" -u "${MYSQL_ADMIN_USERNAME}" -p${MYSQL_ADMIN_PASSWORD} -B -N -e "use \`${REQUIRED_ENV_VARS[DB_DATABASE]}\`; SELECT User FROM mysql.user;")
} &> /dev/null

if [ ${MYSQL_USER_EXISTS} == true ]; then
    printf "${GREEN}User verified!${NC} -- Verified user ${GREEN}${REQUIRED_ENV_VARS[DB_USERNAME]}${NC} was created on database ${GREEN}${REQUIRED_ENV_VARS[DB_DATABASE]}${NC} on host ${GREEN}${REQUIRED_ENV_VARS[DB_HOST]}${NC}\n"
else
    printf "${RED}Warning: User could not be verified!${NC} -- Unable to verify if ${RED}${REQUIRED_ENV_VARS[DB_USERNAME]}${NC} was created on database ${RED}${REQUIRED_ENV_VARS[DB_DATABASE]}${NC} on host ${RED}${REQUIRED_ENV_VARS[DB_HOST]}${NC}\n"
fi