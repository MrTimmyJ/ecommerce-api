let api = '/ecommerce-api/api';
let cart = [];

// Render cart items
function renderCart() {
    const cartItems = document.getElementById("cart-items");
    if (!cartItems) return;

    cartItems.innerHTML = ''; // Clear existing items
    cart.forEach(item => {
        const li = document.createElement("li");
        li.innerHTML = `
            ${item.name} - $${item.price.toFixed(2)} x 
            <input type="number" min="1" value="${item.quantity}" 
                   onchange="updateQuantity(${item.id}, this.value)">
            <button onclick="removeItem(${item.id})">Remove</button>
        `;
        cartItems.appendChild(li);
    });

    updateCartSummary();
}

function renderProduce() {
    const grids = document.getElementsByClassName("product-grid");
    if (grids.length < 1) {
        return;
    }
     
    let endpoint = `${api}/products`
    const grid = grids[0];
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    const query = urlParams.get('query');
    if (query != "") {
      endpoint += '/search' + queryString;
    }
    fetch(endpoint, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' }
    }).then(response => {
        if (!response.ok) {
            throw new Error("Error finding products. Please try again.");
        }
        return response.json();
    }).then(data => {
        grid.innerHTML = "";
        for (const product of data.Products) {
            const card = document.createElement("div");
            card.classList.add("product-card");

            const prod_img = document.createElement("img");
            prod_img.src = "static/img/products/" + product.image;
            prod_img.alt = product.name;
            card.appendChild(prod_img);

            const card_details = document.createElement("div");
            card_details.classList.add("product-details");
            

            const prod_name = document.createElement("h3");
            prod_name.classList.add("product-name");
            prod_name.innerText = product.name;
            card_details.appendChild(prod_name);
            /*
	    const prod_unit = document.createElement("h3");
            prod_unit.classList.add("product-unit");
            prod_unit.innerText = product.unit;
            card_details.appendChild(prod_unit);
            */
	    card.appendChild(card_details);

            const card_checkout = document.createElement("div");
            card_checkout.classList.add("product-card-checkout");

            const prod_price = document.createElement("p");
            prod_price.classList.add("price");
            prod_price.innerText = "$" + (product.price / 100).toFixed(2);
            card_checkout.appendChild(prod_price);

            const prod_id = document.createElement("input");
            prod_id.type = "hidden";
            prod_id.name = "prod_id";
            prod_id.value = product.ID;
            card_checkout.appendChild(prod_id);

            const prod_button = document.createElement("button");
            prod_button.classList.add("add-to-cart-button");
            prod_button.textContent = "Add to Cart";
            card_checkout.appendChild(prod_button);
            card.appendChild(card_checkout);

            grid.appendChild(card);

            prod_button.addEventListener("click", addToCart);
        }
        const loader = grid.previousElementSibling;
        loader.classList.add('hidden');
        grid.classList.remove('hidden');
    }).catch(error => {
        console.error('Error:', error);
        alert(error.message);
    });
}

// Add item to cart using 'e' as a Button Click Event
function addToCart(e) {
    const id = Number(e.target.parentElement.querySelector("input[name='prod_id']").value);
    const originalColor = e.target.style.backgroundColor;
    const flashColor = "#3f5701";
    let tempCart = localStorage.getItem("cart");
    if (tempCart == null || tempCart.length == 0) {
        tempCart = [];
    } else {
        tempCart = JSON.parse(tempCart);
        for (const item of tempCart) {
            if (item.id == id) {
                item.quantity++;
                localStorage.setItem("cart", JSON.stringify(tempCart));
                cart = tempCart;
                e.target.style.backgroundColor = flashColor;
                setTimeout(function (){
                    e.target.style.backgroundColor = originalColor;
                }, 300);
                return;
            }
        }
    }

    let name = e.target.parentElement.parentElement.querySelector(".product-name").innerText;
    let price = e.target.parentElement.querySelector(".price").innerText;
    if (price.charAt(0) == "$") {
        price = price.substring(1);
    }
    price = Number(price);

    tempCart.push({"id": id, "name": name, "price": price, "quantity": 1});
    localStorage.setItem("cart", JSON.stringify(tempCart));
    cart = tempCart;
    e.target.style.backgroundColor = flashColor;
    setTimeout(function (){
        e.target.style.backgroundColor = originalColor;
    }, 300);
}

// Update item quantity
function updateQuantity(itemId, newQuantity) {
    const item = cart.find(product => product.id === itemId);
    if (item) {
        item.quantity = Math.max(parseInt(newQuantity) || 1, 1);
        updateCartSummary();
        localStorage.setItem("cart", JSON.stringify(cart));
        renderCart();
    }
}

// Remove an item from the cart
function removeItem(itemId) {
    cart = cart.filter(product => product.id !== itemId);
    localStorage.setItem("cart", JSON.stringify(cart));
    renderCart();
}

// Update cart summary
function updateCartSummary() {
    const totalItems = cart.reduce((sum, item) => sum + item.quantity, 0);
    const totalPrice = cart.reduce((sum, item) => sum + item.price * item.quantity, 0);
    document.getElementById("cart-count").textContent = totalItems;
    document.getElementById("cart-total").textContent = totalPrice.toFixed(2);
}

// Checkout function
function checkout() {
    // Retrieve cart total
    const total = cart.reduce((sum, item) => sum + item.price * item.quantity, 0).toFixed(2);

    if (cart.length === 0) {
        alert("Your cart is empty!");
        return;
    }

    // Capture user information
    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;
    const address1 = document.getElementById('address1').value;
    const address2 = document.getElementById('address2').value;
    const city = document.getElementById('city').value;
    const state = document.getElementById('state').value;
    const zip = document.getElementById('zip').value;
    //const card = document.getElementById('card').value;
    //const expiry = document.getElementById('expiry').value;
    //const cvv = document.getElementById('cvv').value;

    // Basic validation
    if (!name || !email || !address1 || !city || !state || !zip) {
        alert("Please fill out all fields.");
        return;
    }

    /*
    if (!confirm(`Proceed to checkout with a total of $${total}?`)) {
        return; // Cancel checkout if user declines
    }
    */

    // Prepare payload with user info and cart data
    const payload = {
	    name: name,
      email: email,
	    address_one: address1,
	    address_two: address2,
	    city: city,
      state: state,
      zip: zip,
      items: cart.map(item => ({
          product_id: Number(item.id),
          quantity: Number(item.quantity)
      }))
    };

    // Send data to API
    fetch(`${api}/orders/new`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload)
    }).then(response => {
        if (!response.ok) {
	    console.log(response);
            throw new Error("Error placing order. Please try again.");
        }
        return response.json();
    }).then(data => {
        alert(`Order placed successfully!`);
        cart = []; // Clear the cart on success
        localStorage.setItem("cart", JSON.stringify(cart));
        renderCart();
        document.getElementById('name').value = "";
        document.getElementById('email').value = "";
        document.getElementById('address1').value = "";
        document.getElementById('address2').value = "";
        document.getElementById('city').value = "";
        document.getElementById('state').value = "";
        document.getElementById('zip').value = "";
    }).catch(error => {
        console.error('Error:', error);
        alert(error.message);
    });
}

function search(e, query) {
    if (!query) {
      let el = e.target;
      if (el.tagName == "INPUT") {
        if (e.keyCode != 13) {
          return;
        }
      } else if (el.tagName == "I") {
        el = el.parentElement.previousElementSibling;
      } else {
        el = el.previousElementSibling;
      }
      const searchInput = el;
      query = searchInput.value;
    }
    if (query == "") {
      return;
    }
    const params = new URLSearchParams();
    params.append('query', query);
    location.href = "produce.html?" + params.toString();
}

// Initialize cart on page load
window.onload = renderCart;
window.addEventListener("load", function(){
    let tempCart = localStorage.getItem("cart");
    if (tempCart == null || tempCart.length == 0) {
        cart = [];
    } else {
        cart = JSON.parse(tempCart);
    }
    renderCart();
    renderProduce();
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    const query = urlParams.get('query');
    if (query != "") {
      let searchboxes = document.getElementsByClassName('search-input');
      for (const box of searchboxes) {
        box.value = query;
      }
    }
});
