import React from 'react';
import { connect } from 'react-redux';
import { string, number } from 'prop-types';
import marked from 'marked';
import styled from 'styled-components';

const ProblemLink = styled.h1`
  margin: 10px 0 0 0;
  font-weight: bold;
  font-family: Menlo, monospace;
  color: rgb(204, 65, 65);

  & a {
    font-weight: bold;
    font-family: Menlo, monospace;
    color: #d0021b;
    text-decoration: none;
  }
`;

const Detail = ({ id, title, slug, difficulty, code, explanation, algorithm }) => {
  return (
    <div>
      <ProblemLink>
        <a
          target="_blank"
          href={ `https://leetcode.com/problems/${slug}/` }
        >
          { `${id}. ${title}` }
        </a>
      </ProblemLink>

      <h2>题意</h2>
      <div dangerouslySetInnerHTML={ { __html: marked(explanation) } } />

      <h2>解答</h2>
      <div dangerouslySetInnerHTML={ { __html: marked(algorithm) } } />

      <h2>代码</h2>
      <pre>
        <code style={ {
          borderRadius : '10px',
          fontSize     : '13px',
          lineHeight   : '18px'
        } } lang="golang">{ code }</code>
      </pre>
    </div>
  );
};

Detail.propTypes = {
  id          : string.isRequired,
  title       : string.isRequired,
  slug        : string.isRequired,
  difficulty  : number.isRequired,
  code        : string.isRequired,
  explanation : string.isRequired,
  algorithm   : string.isRequired
};

const mapStateToProps = state => state;

export default connect(mapStateToProps)(Detail);
