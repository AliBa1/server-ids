# Server Intrustion Detection System (IDS)

[![Coverage](https://img.shields.io/badge/Coverage-77.4%25-brightgreen)](https://github.com/AliBa1/server-ids/actions)


## Abstract
This is a Server Intrusion Detection System prototype for a REST API made in Golang. The system detects attacks such as SQL Injection, XSS, broken access control, and potentially more going forward such as login and DoS attacks. The site is hosted and can be accessed [through this link](http://server-ids.up.railway.app). The site serves as a company website for internal members to view shared documents. This shows a real world example of how an IDS can be used to find out about potential attacks on an application that needs to be secure. Going forward an IPS (intrusion prevention system) can be added to it to prevent and stop attacks once they are detected.

## Introduction

When I was first tasked with making an IDS I was thinking of just making something simple that uses Linux commands to detect a few common attacks and call it a day. However, as I thought about it more I wanted to implement the system into something I am interested in to turn this into a more useful learning opportunity overall and something I could possibly use or improve on in the future. I decided to make the IDS for a REST API since I enjoy doing software development and wanted to improve my backend development skills. I made it in Go since it's a language I like but I haven't had the opportunity to use it much up until this point. I chose a REST API as the target because of the variety of security risks that come with it such as SQL injections, XSS attacks, broken access control, login attacks, and DoS attacks. My goal is to develop a system that can detect these attacks effectively before they cause any damage.

Before starting development of the system I looked at it from a high level perspective to see exactly how it should be setup. The way I saw it was that the attacker would make calls to the API through a website then the incoming traffic would be inspected in the middleware before it can be used to do any operations. From there I would check for intrusions and send out alerts if anything is detected. In the case of a detected attack I would alert the admin of what happened and log them in the console. I also may expand it to prevent attacks by setting up firewalls based on which attack got detected to keep the server running without any issues. On top of that I will look to optimize the performance of the IDS by only running the necessary detections and possibly using concurrency or parallelism to run the detection more efficiently since performance is a common concern with intrusion detection systems.

## System Architecture

> **UML (a bit outdated but still close to current structure)** <br> ![UML for the system code](images/Server%20IDS%20UML.drawio.png)

> **High Level Design** <br> ![UML for the system code](images/Server%20IDS%20High%20Level.png)

> **System Process** ![UML for the system code](images/Server%20IDS%20High%20Level%20Complete.png)

## Threat Model

The attacker in this instance would be someone trying to access information or documents that they don't have access to. The attacker could also be trying to alter or destroy information in the database. This attempts to mock a real world example where a company has a site for sharing documents with internal members. The attacker would launch attacks by trying to inject SQL commands, malicious code, or getting access to higher roles to get what they can out of it. The types of attacks they would use are SQL injection, XSS, broken access control, and login attacks. Upon opening the site a login page will be shown. The attacker can use this as a way to launch many different attacks. A SQL injection attack can take place by entering a username or password that would trick the SQL command into doing something it's not supposed to like logging in without the right credentials or destroying a table in the database. By doing a similar action a XSS attack could also take place by inserting code into the URL or into inputs for username or password. This can be done using HTML script tags which the attacker would use by sending out URLs with these tags in them that would run malicious code on the site or on a users device if they click the link. Login attacks could also occur from this page through trying many username and password combinations. Another place where attacks can happen is on the users page. By default the users page cannot be accessed if the user isn't logged in. If a user is logged in they can see all of the users that have access to the site and their roles. If they’re an admin they can change anyones role, otherwise they can't. In this case broken access control can take place since the attack would try to elevate their role to access this page. The attacker could also be someone within the company attempting to elevate their role in order to gain access to more privileges. To change a users’ role from this page there are drop-downs to select their role. If the user doesn’t have access to change roles the drop-down would be disabled. To bypass this an attacker could inspect element to enable the dropdown and change the role anyways. 

![Users page for a guest user](images/users-page.png)
*The user page from the perspective of a user with the guest role*

## Defense Model

To come up with attacks that could possibly occur I checked out the [OWASP Top Ten](https://owasp.org/www-project-top-ten) which shows the most common web application security risks. The defense model is set up to detect most of the attacks that the attacker could use. The types of attacks that the defense model tries to detect are SQL injection, XSS, broken access control, login attacks (haven't implemented yet), and potentially DoS attacks. To do this the IDS is placed in the middleware so that incoming traffic can be inspected before any action is taken on it. The IDS is setup so that it will only run specific detections if they are possible which is implemented using the composite design pattern.

![IDS Function in the middleware](images/ids-function.png)

For detecting SQL injection, I used a signature based detection to identify possible attacks. The rules included using quotes, commands (like DROP, SLEEP, DELETE), and common symbols used in SQL injections (like =, ;, and --). 

![SQL Injection Rules](images/sql-rules.png)

While the rules cover many cases it also will lead to false alarms since people use some of these symbols especially in passwords. To offset this issue I set the alerts as medium severity and made sure the rules are as specific as possible. I check for SQL injection inside of the URL, cookies, HTTP header, and HTTP body of any request with user input. To detect XSS attacks I did something similar using signature based detection as well but only when there are two angle brackets (< and >) in a string. This is because XSS attacks are launched using the script tag in HTML so it would need to have both those brackets to launch an attack. I could have also just made the script tag as the rule but I did angle brackets just incase if they try anything clever using any other HTML tags. I also added the URL encoding (%3C and %3E) for angle brackets as a rule since it can be ran through a URL as well. I check for XSS attacks in the URL, cookies, HTTP header, and HTTP body.

![XSS detection rules](images/check-xss.png)

![XSS sequence](images/xss-sequence.png)

For broken access control I want to check for anytime a role is being elevated using anomaly based detection since admins are the only ones that have the ability to change a users' role. I would check to see if the person changing someones role is logged in or an admin. If they aren’t a user it detects it as an attack with a high severity since it means that they were able to bypass into an protected page before they did the attack. This detection could also extend to someone viewing a locked document without having the correct role or accessing the API with suspicious header information. 

![Broken Access Control code](images/bac-checker.png)

As for login attacks, the detection isn't set up yet but it would check for brute force attempts, suspicious username/password combos, trying to login to many different accounts from the same IP, or login attempts on a users’ account from a suspicious location. This would mainly use anomaly based detection mixed in with some signature based detection as well. For DoS attacks detection isn't set up yet either but it would be based on the amount of requests made in a certain amount of time. Once an attack is detected an alert is sent containing the signature ID, revision number, severity, attack type, date/time, and a message of what happened. The alerts are currently being sent to the console but they can easily be sent as an email, text message, or logged to a file depending on the severity of the attack.. 

![Alert Function](images/alert-admin.png)

To block attacks, requests can be dropped as soon as it is detected since the detector runs in the middleware so the system wouldn't be affected. Some attacks may not be able to be stopped such as a XSS attack where a user clicks a bad link and it takes their cookies since that would be before they access the system but even in this case, using CSRF (cross site request forgery) tokens would be a way to prevent this. Adding a IPS to the IDS would help stop attacks by taking actions like setting up firewalls once attacks are detected. DDoS attacks would be a challenge since I wouldn't be able to use a singular IP to prevent or stop it so the service would likely go down.
