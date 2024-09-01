Project Overview

This repository contains the codebase for an Information Security Project that focuses on secure communication between students and professors, leveraging encryption, key exchange, and digital certificates. The project is divided into several components, including the professor's code, user code, server code, and Certificate Authority (CA) system.
Key Features

    Secure Communication: Establishes secure channels using TCP connections, with encryption protocols like AES and RSA to ensure data confidentiality and integrity.
    User Authentication: Supports secure login and signup processes with encryption for sensitive information.
    Chat Functionality: Facilitates encrypted chat between students and professors, with key pair generation and session key exchange.
    Certificate Management: Implements a Certificate Authority (CA) for generating, signing, and managing digital certificates, crucial for secure communications.
    File Encryption and Signing: Provides functionality to encrypt and sign files, ensuring authenticity and security during transmission.

Project Components

    Professor's Code: Manages the professor's interactions, including login, key exchange, secure chat, and file handling.
    User's Code: Handles user interactions such as signup, login, chat, and secure communication with the server.
    Server Code: Manages the backend operations, including user authentication, chat management, and handling secure communication channels.
    Certificate Authority (CA) Code: Responsible for generating and verifying digital certificates, ensuring secure communication within the system.
    Encryption Package: A set of utilities for AES and RSA encryption/decryption, and digital signature generation/verification.

How to Use

    Clone the Repository:

    bash


    Run the CA Server: Start the CA server to manage certificate requests and issuance.
    Deploy the Server: Start the TCP server that handles student and professor logins and secure communications.
    Execute User/Professor Code: Run the respective client code to interact with the server, performing actions like secure login, chat, and file exchanges.

Security Measures

    Encryption: Utilizes AES-256 CBC for encrypting sensitive data and RSA for secure key exchanges.
    Digital Certificates: Ensures secure communication by verifying client certificates through a CA.
    Data Integrity: Files and messages are encrypted and signed to prevent unauthorized access and tampering.
