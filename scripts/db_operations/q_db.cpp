#include <iostream>
#include <mysql/mysql.h>
#include <string>

const char* server = "localhost";  
const char* user = "root";  
const char* password = "11111111";  
const char* database = "DY"; 

//g++ -std=c++11 q_db.cpp -o q_db -I/opt/homebrew/opt/mysql/include/mysql -L/opt/homebrew/opt/mysql/lib -lmysqlclient

void executeQuery(MYSQL* mysql, const std::string& query) {
    if (mysql_query(mysql, query.c_str())) {
        std::cerr << "Failed to execute query: " << mysql_error(mysql) << std::endl;
        return;
    }

    MYSQL_RES* result = mysql_store_result(mysql);
    if (result) {
        int num_fields = mysql_num_fields(result);
        MYSQL_ROW row;

        while ((row = mysql_fetch_row(result))) {
            for (int i = 0; i < num_fields; i++) {
                std::cout << (row[i] ? row[i] : "NULL") << " | ";
            }
            std::cout << std::endl;
        }

        mysql_free_result(result);
    } else {
        if (mysql_field_count(mysql) == 0) {
            std::cout << "Query OK, " << mysql_affected_rows(mysql) << " rows affected" << std::endl;
        } else {
            std::cerr << "Failed to store result set: " << mysql_error(mysql) << std::endl;
        }
    }
}

int main(int argc, char* argv[]) {
    if (argc != 2) {
        std::cerr << "Usage: " << argv[0] << " [-1 | -2 | -3 | -4 | -5 | -qa]" << std::endl;
        return 1;
    }

    std::string param = argv[1];
    std::string query;
    std::string tableName;

    if (param == "-1") {
        query = "SELECT * FROM user_login_info;";
        tableName = "user_login_info";
    }else if (param == "-2") {
        query = "SELECT * FROM user;";
        tableName = "user";
    }else if (param == "-3") {
        query = "SELECT * FROM video_records;";
        tableName = "video_records";
    }else if (param == "-4") {
        query = "SELECT * FROM favorite;";
        tableName = "favorite";
    }else if (param == "-5") {
        query = "SELECT * FROM comment;";
        tableName = "comment";
    }else if (param == "-qa") {
        query = "SHOW TABLES;";
        tableName = "All Tables";
    } else {
        std::cerr << "Invalid parameter"<< std::endl;
        std::cerr << "Use -1 for user_login_info" << std::endl;
        std::cerr << "Use -2 for user" << std::endl;
        std::cerr << "Use -3 for video_records" << std::endl;
        std::cerr << "Use -4 for favorite" << std::endl;
        std::cerr << "Use -5 for comment" << std::endl;
        std::cerr << "Use -qa to query all tables." << std::endl;
        return 1;
    }

    MYSQL* mysql = mysql_init(nullptr);
    if (!mysql) {
        std::cerr << "Failed to initialize MySQL connection: " << mysql_error(mysql) << std::endl;
        return 1;
    }

    if (!mysql_real_connect(mysql, server, user, password, database, 0, nullptr, 0)) {
        std::cerr << "Failed to connect to database: " << mysql_error(mysql) << std::endl;
        mysql_close(mysql);
        return 1;
    } else {
        std::cout << "Connected to MySQL server, querying table: " << tableName << std::endl;
    }
    executeQuery(mysql, query);

    mysql_close(mysql);
    return 0;
}