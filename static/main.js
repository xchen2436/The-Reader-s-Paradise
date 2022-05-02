let carts = document.querySelectorAll('.add-cart');
let products = [
{
    name:'THE ALCHEMIST PAULO COELHO',
    tag:'b1',
    price: 25,
    inCart: 0
},
{
    name:'ONLY LOVE IS REAL',
    tag:'b2',
    price: 22,
    inCart: 0
},
{
    name:'HARRY POTTER and the philosophers Stone',
    tag:'b3',
    price: 30,
    inCart: 0
},
{
    name:'LOOKING FOR ALASKA',
    tag:'b4',
    price: 20,
    inCart: 0

},
]

for(let i=0;i < carts.length; i++){
    carts[i].addEventListener('click',()=>{
        cartNumbers(products[i])
        totalCost(products[i])
    })
}
 function cartNumbers(products){
     let productNumbers = localStorage.getItem('cartNumbers');
     productNumbers = parseInt(productNumbers);
     if (productNumbers) {
        localStorage.setItem('cartNumbers',productNumbers+ 1);
     }else{
        localStorage.setItem('cartNumbers', 1);
     }
     setItems(products);
 
 }

function setItems(product){
    let carItems = localStorage.getItem('ProductsInCart');
    carItems = JSON.parse(carItems);
    if(carItems != null){
        if(carItems[product.tag] == undefined ){
            carItems = {
                ...carItems,
                [product.tag]:product
            }
        }
        carItems[product.tag].inCart += 1;
    }else{
        product.inCart = 1;
        carItems = {
            [product.tag]:product
        }
    }
    localStorage.setItem("ProductsInCart",JSON.stringify(carItems));
}


function totalCost(product){
    let cartCost = localStorage.getItem('totalCost')

    if(cartCost != null){
        cartCost = parseInt(cartCost);
        localStorage.setItem("totalCost",cartCost + product.price);
    }else{
        localStorage.setItem("totalCost",product.price)
    }

}

function displayCart(){
let carItems = localStorage.getItem("ProductsInCart");
let cartCost = localStorage.getItem('totalCost')
carItems = JSON.parse(carItems);
let productContainer = document.querySelector(".products")
if(carItems && productContainer){
    //console.log("run")
    productContainer.innerHTML = '';
    Object.values(carItems).map(item => {
        productContainer.innerHTML += `
        <table>
            <td class="cartproduct" ><img src="../static/images/${item.tag}.jpg" style="height:40px;width:30px;"> </td>
            <td class="carttitle">${item.name} </td>
            <td class="cartprice">$${item.price}.00 </td>
            <td class="cartquantity">${item.inCart} </td>   
            <td class="carttotal">$${item.inCart * item.price}.00</td>
        </table>
        `;
    });
    productContainer.innerHTML += `
        <div class = "TotalPrice">
            <h4 class="baskTitle"> Total Price</h4>
            <h4 class="baskTotal"> $${cartCost},00</h4>
        </div>
        <div> 
        <button id="go" class="buy" onclick="buy()">Purchase</button>
        <button id="go" class="buy" onclick="removeitem()">Clear All</button>
        </div>
    `;
}

}
function removeitem(){
    window.localStorage.removeItem('ProductsInCart');
    window.localStorage.removeItem('totalCost');
    window.localStorage.removeItem('cartNumbers');
    window.location.reload();
}

function buy(){
    let cartCost = localStorage.getItem('totalCost')
    alert("Payment successful, the total price is $" + cartCost + ",00");
}
displayCart()

