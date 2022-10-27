import socket
import random
import time

def dropOrDelay(number) :
    """
    drop : 10 % of probability to return true
    delay : delay 100ms

    return True : 
        if number is even , drop the packet.
        if number is odd , delay 100ms to packet.

    return False : 
        do nothing
    """
    if  number % 2 == 0 :
        return True if random.randint(0,9) == 0 else False
    else :
        time.sleep(0.1)
        return True if random.randint(0,19) == 0 else False

localIP     = "140.118.122.155"
localPort   = 5406

# Create a datagram socket
UDPProxySocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)

# Bind to address and ip
UDPProxySocket.bind((localIP, localPort))

print("UDP Proxy up and listening")

# Listen for incoming datagrams
while(True):
    proxyMsg , proxyIP = UDPProxySocket.recvfrom(1024)

    print("Message from Client: ",proxyMsg.decode())
    print("Client IP Address: ",proxyIP)

    # Sending a msg to server
    UDPProxySocket.sendto(proxyMsg, ("140.118.122.155", 5405))