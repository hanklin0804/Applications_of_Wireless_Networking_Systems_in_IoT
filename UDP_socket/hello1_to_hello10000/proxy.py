import socket

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