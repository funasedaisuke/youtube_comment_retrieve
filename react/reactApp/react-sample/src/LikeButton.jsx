import React from "react";

const LikeButton = (props) =>{
    return (
        <div>
            <button id="counter">いいね数:{props.count}</button>
        </div>
    )

}
export default LikeButton