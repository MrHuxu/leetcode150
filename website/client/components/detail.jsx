import React from 'react';
import { connect } from 'react-redux';
import { string } from 'prop-types';
import marked from 'marked';
import styled from 'styled-components';

const Container = styled.div`
  background-color: #eee;
`;

const Detail = ({ solution, explanation }) => {
  return (
    <Container>
      <pre><code lang="golang">{ solution }</code></pre>
      <div dangerouslySetInnerHTML={ { __html: marked(explanation) } } />
    </Container>
  );
};

Detail.propTypes = {
  solution    : string.isRequired,
  explanation : string.isRequired
};

const mapStateToProps = state => state;

export default connect(mapStateToProps)(Detail);
