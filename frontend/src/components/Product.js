import React from "react";

export default function Product(props) {

    const { product, onAdd } = props;

    return (
        <div>
            <img className="small" src={product.Image} alt={product.Name}></img>
            <h3>{product.Name}</h3>
            <div>${product.Price}</div>
            <div>
                <button onClick={() => onAdd(product)}>Add to cart</button>
            </div>
        </div>
    )
}