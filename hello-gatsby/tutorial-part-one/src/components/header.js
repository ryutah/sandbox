import React from 'react';
import PropTypes from 'prop-types';

const Header = ({ headerText }) => <h1>{headerText}</h1>;

Header.propTypes = {
  headerText: PropTypes.string.isRequired,
};

export default Header;
