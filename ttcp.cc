#include <iostream>
#include "inet_address.h"

int main() {

    InetAddress addr(9090, true);
    InetAddress cpaddr(addr.getSockAddr());
    std::cout << addr.toEndpoint() << std::endl;
    std::cout << cpaddr.toEndpoint() << std::endl;
}