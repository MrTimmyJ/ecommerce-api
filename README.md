# The Evergreen State College E-commerce Web App
E-commerce project for The Evergreen State College 2024 DSA Group Project. Used to practice sprints and standup for development.

Author: Team of 8 - Timothy Johnson(Full-Stack) <br>
Date: October 2024 to December 2024

### 🔗 Live Demo

[The Evergreen State College Farm E-Commerce Website](https://www.tesc.farm/)

### 🔗 Original Github:

[The Evergreen State College Farm E-Commerce Github](https://github.com/Acstrayer/TESCSE-Ecom)

## Overview

TESCSE-Ecom simulates a small online produce store, allowing users to:

    Browse seasonal fruits and vegetables

    Search and filter items

    Create user accounts and place orders

    View shopping carts

    Manage data from a secure admin dashboard

It served as a capstone for group collaboration, sprint development, and full-stack architecture using Go and modern web technologies.

⚙️ Features

  🛒 User-Facing

    ✅ Responsive product grid with seasonal highlights

    🔍 Live search functionality

    🛍️ Shopping cart interface

🧑‍💼 Admin-Facing

    👤 Employee secure login

    📦 View and manage all products and orders

    📈 Dashboard with profit graph

    ➕ Add/update/delete produce via REST API

💻 Tech Stack

| Layer         | Technology                 |
| ------------- | -------------------------- |
| Backend API   | Go (Gin framework)         |
| Database      | SQLite                     |
| Frontend      | HTML, CSS, JS              |
| UI Framework  | Bootstrap                  |
| Auth          | Session-based auth         |
| Visualization | Chart.js (Admin Dashboard) |

🧩 Architecture

    RESTful API using Gin

    Models and controllers for products, orders, and users

    JSON-based communication with the frontend

    Modular layout with <header>, <main>, and <footer>

    Dynamic elements like seasonal spotlight, interactive search bar, and cart

📁 Code Structure

. <br>
TESCSE-Ecom/ <br>
├── api/ &nbsp;&nbsp;&nbsp;---&nbsp;&nbsp;&nbsp; Go backend <br>
│   ├── main.mod <br>
│   ├── main.sum <br>
│   ├── main.go <br>
│   ├── userstable.sql <br>
│   ├── controllers/ <br>
│   └── models/ <br>
├── public_html/ <br> &nbsp;&nbsp;&nbsp;---&nbsp;&nbsp;&nbsp; HTML / CSS frontend 
│   ├── index.html <br>
│   ├── produce.html <br>
│   ├── cart.html <br>
│   ├── about.html <br>
│   ├── contact.html <br>
│   ├── signin.html <br>
│   ├── dashboard.html <br> &nbsp;&nbsp;&nbsp;---&nbsp;&nbsp;&nbsp; Bootstrap dashboard 
│   └── static/ <br>
│       ├── img/ <br>
│       ├── style.css <br>
│       └── script.js <br>

🖼️ Screenshots / Visuals

![tescfbanner](https://github.com/user-attachments/assets/baad549c-077b-4286-968e-9087af4a7514)

🚀 Getting Started

    1. Clone the repository:

      git clone https://github.com/Acstrayer/TESCSE-Ecom

     2. Backend:
      
      cd TESCSE-Ecom/api
      go run main.go

    3. Frontend:

      cd TESCSE-Ecom/static
      open index.html
    

🪪 License

This open-source project is available under the [Creative Commons Zero v1.0 Universal (CC0-1.0) License](https://creativecommons.org/publicdomain/zero/1.0/).

[//]: # (Test Jenkins CI/CD - Mon Nov 17 08:39:07 UTC 2025)

