import socket
import random
import time
        

localIP     = "140.118.122.155"
localPort   = 5406

# Create a datagram socket
UDPProxySocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)

# Bind to address and ip
UDPProxySocket.bind((localIP, localPort))

print("UDP Proxy up and listening")

# Listen for incoming datagrams
drop_num = 0
deley_num = 0
while(True):
    try: 
        proxyMsg , proxyIP = UDPProxySocket.recvfrom(1024)

        # print("Message from Client: ",proxyMsg.decode())
        # print("Client IP Address: ",proxyIP)

        # it drops each received packet with 10% probability if i is even number. 
        if int(proxyMsg.decode().split(' ')[1]) % 2 == 0:
            if not random.randint(0,9) == 0:
                # Sending a msg to server
                UDPProxySocket.sendto(proxyMsg, ("140.118.122.155", 5405))
            else: 
                print(proxyMsg.decode(),'is dropped')
                drop_num += 1 

        # it delays 100 ms the received packet with 5% probability before forwarding to the server if i is odd number. 
        else:
            if random.randint(0,19) == 0: 
                time.sleep(0.1)
                deley_num += 1
                print(proxyMsg.decode(),' id delays 100 ms ')
            # Sending a msg to server
            UDPProxySocket.sendto(proxyMsg, ("140.118.122.155", 5405))
        if int(proxyMsg.decode().split(' ')[1]) >= 10000:
            print('num of drop: ',drop_num)
            print('num of delay: ',deley_num)
            break
    except:
        # Create a datagram socket
        UDPProxySocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)

        # Bind to address and ip
        UDPProxySocket.bind((localIP, localPort))

        
