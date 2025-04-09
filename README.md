# Network Intrustion Detection System (IDS)

[![Coverage](https://img.shields.io/badge/Coverage-75.1%25-yellow)](https://github.com/AliBa1/server-ids/actions)

Network Intrusion Detection System prototype for a REST API using HTTP

## Abstract

## Introduction

When I was first tasked with making a network intrusion detection system (IDS)
I was initally thinking of just making something simple using Linux commands
to detect a few coommon attacks and call it a day. However, as I thought about it
more I wanted to implement the system into something I am interested in to turn this into
a valuable learning opportunity and something I could possibly use in the future. I
decided to make the IDS for a REST API since I enjoy doing full-stack development
and wanted to improve my backend development skills. I made it in Go since it's a language
I enjoy but haven't got to use much yet. I chose a REST API as the target because of the variety
of security risks that come with it such as SQL injections, XSS attacks, and broken access control.
My goal is to develop a system that can detect these attacks effectively before they cause any damage.

Before starting development of the system I looked at it from a high level perspective
to see exaclty how it should be made. The way I saw it was that the attacker would make a call to the
API then the incoming traffic would be inspected in the middleware before it can be recieved.
From there I would check for intrusions and decide to pass or not. From here I would alert the admin of any
detections and log them to a file for future reference. I also may expand it to prevent attacks by setting up
firewalls based on what attack is detected to keep the server running wihtout any issues.
On top of that I may also look to optimize the performance of the detection system by implementing the
composite design pattern to run checks for only for attacks that are possible with the incoming traffic.
I also may use multithreading to run the system faster which would allow the server to perform as best as it can.

## System Architecture (_will include network setup_)

> **UML** <br> ![UML for the system code](images/Server%20IDS%20UML.drawio.png)

> **High Level Design** ![UML for the system code](images/Server%20IDS%20High%20Level.png)

> **System Process** ![UML for the system code](images/Server%20IDS%20High%20Level%20Complete.png)

## Threat Model

## Defense Model
