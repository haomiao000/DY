#include <mysql/mysql.h>
#include <iostream>
#include <string>
#include <vector>


const char* server = "localhost";  
const char* user = "root";  
const char* password = "11111111";  
const char* database = "DY"; 
//g++ -std=c++11 d_db.cpp -o d_db -I/opt/homebrew/opt/mysql/include/mysql -L/opt/homebrew/opt/mysql/lib -lmysqlclient
// 删除特定表
void executeDeleteTable(MYSQL* conn, const std::string& tableName) {
    std::string query = "DROP TABLE IF EXISTS " + tableName + ";";
    
    if (mysql_query(conn, query.c_str())) {
        std::cerr << "Failed to delete table " << tableName << ": " << mysql_error(conn) << std::endl;
    } else {
        std::cout << "Deleted table successfully: " << tableName << std::endl;
    }
}

// 获取所有表名
std::vector<std::string> getAllTables(MYSQL* conn) {
    std::vector<std::string> tables;
    if (mysql_query(conn, "SHOW TABLES")) {
        std::cerr << "Failed to list tables: " << mysql_error(conn) << std::endl;
        return tables;
    }

    MYSQL_RES* result = mysql_store_result(conn);
    if (result == nullptr) {
        std::cerr << "Failed to store result: " << mysql_error(conn) << std::endl;
        return tables;
    }

    MYSQL_ROW row;
    while ((row = mysql_fetch_row(result))) {
        tables.push_back(row[0]);
    }

    mysql_free_result(result);
    return tables;
}

// 删除所有表
void deleteAllTables(MYSQL* conn) {
    std::vector<std::string> tables = getAllTables(conn);
    for (const auto& table : tables) {
        executeDeleteTable(conn, table);
    }
}

int main(int argc, char* argv[]) {
    if (argc != 2) {
        std::cerr << "Usage: " << argv[0] << " [-1 | -2 | -da]" << std::endl;
        return 1;
    }

    std::string param = argv[1];
    std::string tableName;

    if (param == "-1") {
        tableName = "user_login_info";
    } else if (param == "-2") {
        tableName = "user";
    } else if (param == "-da") {
        tableName = "";
    } else {
        std::cerr << "Invalid parameter. Use -1 for user_login_info, -2 for user and -da to delete all tables." << std::endl;
        return 1;
    }

    MYSQL* conn = mysql_init(nullptr);

    if (conn == nullptr) {
        std::cerr << "mysql_init() failed" << std::endl;
        return EXIT_FAILURE;
    }

    if (mysql_real_connect(conn, server, user, password, database, 0, nullptr, 0) == nullptr) {
        std::cerr << "mysql_real_connect() failed: " << mysql_error(conn) << std::endl;
        mysql_close(conn);
        return EXIT_FAILURE;
    } else {
        std::cout << "Opened database successfully" << std::endl;
    }

    if (param == "-da") {
        std::cout << "Deleting all tables" << std::endl;
        deleteAllTables(conn);
    } else {
        std::cout << "Deleting table: " << tableName << std::endl;
        executeDeleteTable(conn, tableName);
    }

    mysql_close(conn);
    return 0;
}