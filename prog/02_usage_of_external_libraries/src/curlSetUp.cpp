#include "../include/curlSetUp.hpp"

/**
 * @brief Function for curl magic. Writes data from link to string
 */
static size_t writeCallback(void *contents, size_t size, size_t nmemb, void *userp) {
    ((std::string*)userp)->append((char*)contents, size * nmemb);
    return size * nmemb;
}

void curlSetUp(const char url[], std::string *readBuffer) {
    CURL *handle = curl_easy_init();
    if (!handle) {std::cout << "Init failed" << std::endl; return;}

    curl_easy_setopt(handle, CURLOPT_URL, url);
    curl_easy_setopt(handle, CURLOPT_WRITEFUNCTION, writeCallback);
    curl_easy_setopt(handle, CURLOPT_WRITEDATA, readBuffer);

    CURLcode result = curl_easy_perform(handle);
    if (result != CURLE_OK) {std::cout << "Download failed" << std::endl; return;}

    curl_easy_cleanup(handle);

    return;
}