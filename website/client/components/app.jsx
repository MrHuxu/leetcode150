import React from 'react';
import styled from 'styled-components';
import { node } from 'prop-types';

const Container = styled.div`
  max-width: 800px;
  padding: 30px 10px;
  margin-left: calc(50% - 400px);

  @media (max-width: 800px) {
    padding: 10px 10px;
    margin-left: 0;
  }
`;

const App = ({ children }) => {
  return (
    <Container>
      { children }
    </Container>
  );
};

App.propTypes = {
  children : node.isRequired
};

export default App;
