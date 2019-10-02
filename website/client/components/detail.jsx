import React from 'react';
import { connect } from 'react-redux';
import { string, number } from 'prop-types';
import marked from 'marked';
import styled from 'styled-components';

const Container = styled.div`
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
  Oxygen-Sans, Ubuntu, Cantarell, "Helvetica Neue", sans-serif;
  line-height: 1.5;
  font-size: 80%;
  box-sizing: inherit;
  padding: 0 30px;
  box-shadow: inset 580px 0 #fff;
  min-height: 100vh;
  display: block;
`;
const ProblemLink = styled.h1`
  box-sizing: inherit;
  margin: 0;
  font: inherit;
  font-weight: bold;
  font-family: Menlo, monospace;
  padding: 40px 0px 10px 0px;
  color: rgb(204, 65, 65);
  font-size: 2em;
  position: relative;
  clear: both;
  width: 550px;

  & a {
    font: inherit;
    font-weight: bold;
    font-family: Menlo, monospace;
    box-sizing: inherit;
    color: #d0021b;
    text-decoration: none;
  }
`;
const IntentionContainer = styled.div`
  width: 560px;

  @media(max-width: 1200px) {
    width: 100%;
  }
`;
const SolutionTitle = styled.h2`
  font-weight: bold;
  font-size: 1.5em;
  color: rgb(45, 45, 45);
`;
const Right = styled.div`
  position: fixed;
  top: 0;
  right: 0;
  height:100%;
  width: calc(100% - 950px);
  height: 100%;
  overflow: auto;
  background-color: #2d2d2d;
  z-index: -1;

  @media(max-width: 1200px) {
    display: none;
  }
`;
const AlgorithmContainer = styled.div`
  vertical-align: top;
  display: inline-block;
  width: 560px;

  @media(max-width: 1200px) {
    width: 100%;
  }
`;
const CodeContainer = styled.pre`
  vertical-align: top;
  width: 100%;
  margin: 0;

  @media(min-width: 1201px) {
    margin: 0 0 0 10px;
    display: inline-block;
    width: calc(100% - 570px);
  }

  & code {
    padding: 10px 20px;
    background-color: #2d2d2d !important;

    @media(max-width: 1200px) {
      border-radius: 10px;
    }
  }
`;

const Detail = ({ id, title, slug, difficulty, solution, explanation, algorithm }) => {
  return (
    <Container>
      <ProblemLink>
        <a
          target="_blank"
          href={ `https://leetcode.com/problems/${slug}/` }
        >
          { `${id}. ${title}` }
        </a>
      </ProblemLink>

      <IntentionContainer dangerouslySetInnerHTML={ { __html: marked(explanation) } } />

      <Right />

      <SolutionTitle>解答</SolutionTitle>

      <div>
        <AlgorithmContainer dangerouslySetInnerHTML={ { __html: marked(algorithm) } } />

        <CodeContainer>
          <code lang="golang">{ solution }</code>
        </CodeContainer>
      </div>
    </Container>
  );
};

Detail.propTypes = {
  id          : string.isRequired,
  title       : string.isRequired,
  slug        : string.isRequired,
  difficulty  : number.isRequired,
  solution    : string.isRequired,
  explanation : string.isRequired,
  algorithm   : string.isRequired
};

const mapStateToProps = state => state;

export default connect(mapStateToProps)(Detail);
