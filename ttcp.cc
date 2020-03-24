#include <iostream>
#include "accept.h"
#include "inet_address.h"


int main() {
   Acceptor accpetor(InetAddress(19090, false));
   accpetor.accept();
}