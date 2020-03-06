import React from 'react';
import PropTypes from 'prop-types';
import Header from './Header';

const layoutStyle = {
  margin: 20,
  padding: 20,
  border: '1px solid #DDD',
};

const Layout = ({ children }) => (
  <div style={layoutStyle}>
    <Header />
    {children}
  </div>
);

Layout.propTypes = {
  children: PropTypes.node.isRequired,
};

export default Layout;
