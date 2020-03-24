#include "tcp_stream.h"

TcpStream::TcpStream(Socket&& socket)
 : __sock(std::move(socket))
{

}