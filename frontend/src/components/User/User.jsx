import React, { Component } from 'react';

class User extends Component {
  constructor(props) {
    super(props);

    this.state={
      userName: "",
    }
  }

  componentDidMount() {
    const { userId } = this.props.match.params;
    console.log(userId);
    fetch(`http://192.168.1.142:8080/users/${userId}`, {
      credentials: 'include',
    }).then(res => {
      return res.json();
    }).then(data => {
      this.setState({
        userName: data.username
      })
    })
  }
  render() {
    return (
      <div>
        <h1>User Page</h1>
        <h1>Hi I am {this.state.userName}</h1>
      </div>
    );
  }
}

export default User;
