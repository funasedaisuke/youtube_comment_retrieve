// import React from 'react';
// import LikeButton from './LikeButton';

// class App extends React.Component {
//   constructor(props) {
//     super(props);
//     this.state = {
//       count: 0
//     }
//   }

//   componentDidMount() {
//     document.getElementById("counter").addEventListener('click', () => this.setState({count: this.state.count + 1}))
//   }

//   componentDidUpdate(){
//     if (this.state.count >= 5) {
//       this.setState({
//         count: 0
//       })
//     }
//   }
//   render() {
//     return (
//       <div>
//         <LikeButton count={this.state.count} />
//       </div >
//     )
//   }
// }
// export default App

import React, {useState, useEffect} from 'react'

const App = () => {

    const [posts, setPosts] = useState([])

useEffect(() => {
    fetch('http://localhost:8080/index', {method: 'GET'})
    .then(res => res.json())
    .then(data => {
        setPosts(data)
    })
},[])

return (
    <div>
        <ul>
            {
                posts.map(post => <li　key={post.updated_time}>{post.comment}</li>)
            }
        </ul>
        
    </div>
)
}
export default App