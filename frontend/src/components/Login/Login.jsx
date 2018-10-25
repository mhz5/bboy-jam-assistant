import React, { Component } from 'react';
import './Login.scss';
import CredentialsForm from 'components/common/CredentialsForm';


class Login extends Component {
  render() {
    return (
      <CredentialsForm
        action="Log In"
        redirectPath="/users/14"
        submitActionUrl="http://192.168.1.142:8080/login"
      />
    );
  }
}

export default Login;
