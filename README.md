# Middleman-API

This repository contains all the code used by 0xRay's servers for accessing indexed collections, storing anonymous user data, etc.

**Functions**
- Anonymous authentication API based off user hardwareIDs/JWTs
- Serverside caching of authentication credentials/indexed collection data

**Important Functions**
- [Access Contract Data](database/fiberhandlers.go)
- [Authentication](jwt/jwthandler.go)

**Middleman API Diagram**

![Diagram](https://media.discordapp.net/attachments/893372833743372321/1096610569383059506/Blank_diagram_1.png?width=1045&height=538)