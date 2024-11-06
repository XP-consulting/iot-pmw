# IoT-PMW: Secure IoT Passthrough Middleware

This project aims to create an IoT Passthrough Middleware (IoT-PMW) that enables secure data transmission through a completely isolated firewall system. By "complete block system," we mean a firewall that fully restricts all network connections, blocking both inbound and outbound traffic. This middleware is designed to operate on a single server, divided into two separate Docker containers. One container is connected to the external network, while the other connects to the internal network, allowing secure, controlled data exchange with specific JSON protocols across the firewall.

## Why IoT-PMW?

The firewall’s role is to protect the network from external threats by blocking all unauthorized connections. So why introduce a middleware solution? The goal is to enable the safe and selective passage of critical data, specifically JSON protocols, between network segments. Using this system, we can monitor and control network data flows, in our case leveraging Go’s SigML package to track and securely transmit sensor data across network boundaries. This approach ensures secure data handling within a firewall-protected environment.

## Technologies Used

The project primarily utilizes the Go programming language and the SigML package to manage errors and streamline JSON protocol handling. Key tools and methods include:

- **Docker Containers**: Two Docker containers separate trusted and untrusted networks, ensuring secure data flow between global and local networks. Each container maintains a secure connection to prevent man-in-the-middle or other security attacks.
  
- **Encryption**: To further enhance security, encryption is implemented for communication between containers. This encrypted connection aims to prevent unauthorized access to sensitive data.
  
- **Database Integration**: The middleware is designed to be compatible with various databases, but PostgreSQL is our chosen solution for this project.

