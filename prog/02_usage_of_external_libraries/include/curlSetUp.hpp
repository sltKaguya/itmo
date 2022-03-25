#pragma once

#include <curl/curl.h>
#include <iostream>
#include <string>

/**
 * @brief Get the info from given link and write it to the given string
 * 
 * @param url link to get info
 * @param readBuffer pointer to the string to write in
 */
void curlSetUp(const char url[], std::string *readBuffer);