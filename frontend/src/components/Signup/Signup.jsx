import React, { Component } from 'react';
import './Signup.scss';
import CredentialsForm from 'components/common/CredentialsForm';


class Signup extends Component {
  render() {
    return (
      <CredentialsForm
        action="Signup"
        redirectPath="/user/1"
        submitActionUrl="http://localhost:8080/users"
      />
    );
  }
}

export default Signup;
