import socket

localIP     = "140.118.122.155"
localPort   = 5405



# Create a datagram socket
UDPSeverSocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)

# Bind to address and ip
UDPSeverSocket.bind((localIP, localPort))

print("UDP Sever up and listening")

n=0
# Listen for incoming datagrams
while(True):
    SeverMsg, SeverIP = UDPSeverSocket.recvfrom(1024)
    n+=1
    print("Message from Proxy: ",SeverMsg.decode())
    print("Proxy IP Address: ",SeverIP,"\n")

    msg = 'World ' +str(SeverMsg.decode().split(' ')[1])
    # Sending a msg to Client
    UDPSeverSocket.sendto(msg.encode(), ("140.118.122.155", 5407))

