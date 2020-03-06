import React from 'react';
import PropTypes from 'prop-types';
import Container from '../components/container';
import styles from './about-css-modules.module.css';

console.log(styles);

const User = ({ avatar, username, excerpt }) => (
  <div className={styles.user}>
    <img src={avatar} className={styles.avatar} alt="" />
    <div className={styles.description}>
      <h2 className={styles.username}>{username}</h2>
      <p className={styles.excerpt}>{excerpt}</p>
    </div>
  </div>
);

User.propTypes = {
  avatar: PropTypes.string.isRequired,
  username: PropTypes.string.isRequired,
  excerpt: PropTypes.string.isRequired,
};

export default () => (
  <Container>
    <h1>About CSS Modules</h1>
    <p>CSS Modules are cool</p>
    <User
      username="Jane Doe"
      avatar="https://s3.amazonaws.com/uifaces/faces/twitter/adellecharles/128.jpg"
      excerpt="I'm Jane Doe. Lorem ipsum dolor sit amet, consectetur adipisicing elit."
    />
    <User
      username="Bob Smith"
      avatar="https://s3.amazonaws.com/uifaces/faces/twitter/vladarbatov/128.jpg"
      excerpt="I'm Bob Smith, a vertically aligned type of guy. Lorem ipsum dolor sit amet, consectetur adipisicing elit."
    />
  </Container>
);
