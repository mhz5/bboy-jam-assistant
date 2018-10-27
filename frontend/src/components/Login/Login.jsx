import React, { Component } from 'react';
import './Login.scss';
import CredentialsForm from 'components/common/CredentialsForm';


class Login extends Component {
  render() {
    return (
      <CredentialsForm
        action="Log In"
        redirectPath="/"
        submitActionUrl={`${process.env.REACT_APP_API_URL}/login`}
      />
    );
  }
}

export default Login;
