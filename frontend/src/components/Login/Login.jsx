import React, { Component } from 'react';
import './Login.scss';
import CredentialsForm from 'components/common/CredentialsForm';


class Login extends Component {
  render() {
    return (
      <CredentialsForm
        action="Log In"
        submitActionUrl="http://localhost:8080/users"
      />
    );
  }
}

export default Login;
