import socket
import random
import time
from concurrent.futures import ThreadPoolExecutor
import threading
        

localIP     = "127.0.0.1"
localPort   = 5406

# Create a datagram socket
UDPProxySocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)

# Bind to address and ip
UDPProxySocket.bind((localIP, localPort))

print("UDP Proxy up and listening")

# Listen for incoming datagrams
global drop_num 
global deley_num 
global i
drop_num = 0
deley_num = 0
i = 1
# create a lock
lock = threading.Lock()



def proxy():
        print(1)
        # acquire the lock
        lock.acquire()

        proxyMsg , proxyIP = UDPProxySocket.recvfrom(1024)

        # release the lock
        lock.release()
        print(2)
        if int(proxyMsg.decode().split(' ')[1]) > i:
            i = int(proxyMsg.decode().split(' ')[1])
            print(2.5)
            print("Message from Client: ",proxyMsg.decode())
            print("Client IP Address: ",proxyIP)

            # it drops each received packet with 10% probability if i is even number. 
            if int(proxyMsg.decode().split(' ')[1]) % 2 == 0:
                if not random.randint(0,9) == 0:
                    # Sending a msg to server
                    UDPProxySocket.sendto(proxyMsg, ("127.0.0.1", 5405))
                else: 
                    drop_num += 1
                    print(proxyMsg.decode(),'is dropped')


            # it delays 100 ms the received packet with 5% probability before forwarding to the server if i is odd number. 
            else:
                if random.randint(0,19) == 0: 
                    print(proxyMsg.decode(),' id delays 100 ms ')
                    deley_num += 1
                    time.sleep(0.1)
                    
                # Sending a msg to server
                UDPProxySocket.sendto(proxyMsg, ("127.0.0.1", 5405))
            print(3)
        else:
            print("pass")
            pass


            



with ThreadPoolExecutor() as executor:    # 改用 with...as
    while(i<=10000):
        executor.submit(proxy, )
        executor.submit(proxy, )
        executor.submit(proxy, )


print('num of drop: ',drop_num)
print('num of delay: ',deley_num)



        
