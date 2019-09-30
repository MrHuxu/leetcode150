import React from 'react';
import { connect } from 'react-redux';
import { string, number } from 'prop-types';
import marked from 'marked';
import styled from 'styled-components';

import { difficultyContainers, difficultyLabels } from './common';

const Container = styled.div`
  height: 100%;
`;
const Left = styled.div`
  height: 100%;
  width: calc(50% - 60px);
  color: #555;
  padding: 0 0 20px 30px;
  font-size: 12px;
 
  & li {
    line-height: 1.9;
  }
  & blockquote {
    padding: 15px 0 15px 15px;
    margin: 0 0 18px;
    border-left: 5px solid #D1D0CE;
    line-height: 28px;
    font-weight: normal;
    font-size: 12px;
    font-style: italic;
    color: #696969;
  }
  & img {
    max-width: 100%;
  }
  & a {
    color: #4183c4;
  }
  & hr {
    border: 0;
    color: #ddd;
    background-color: #ddd;
    height: 1px;
    margin: 5px 0 19px 0;
  }
  & code {
    display: inline;
    word-wrap: break-word;
    font-size: 12px;
    color: rgb(85, 85, 85);
    background: rgb(255, 255, 255);
    border-width: 1px;
    border-style: solid;
    border-color: rgb(221, 221, 221);
    border-image: initial;
    border-radius: 4px;
    padding: 1px 3px;
    margin: -1px 1px 0px;
  }
`;
const ProblemLink = styled.div`
  font-size: 24px;
  text-align: center;
  padding: 25px 20px 0px 20px;

  & a {
    color: #666;
  }
`;
const Right = styled.div`
  position: fixed;
  top: 0;
  right: 0;
  width: 50%;
  height: 100%;
  overflow: auto;

  & pre code {
    display: block;
    font-size: 0.92rem;
    line-height: 18px;
    font-weight: 12px;
    letter-spacing: 0.03rem;
    margin: 0 0 20px 0;
    padding: 15px !important;
    background-color: #f7f7f7 !important;
    border-width: 0;
  }
`;

const Detail = ({ id, title, slug, difficulty, solution, explanation, algorithm }) => {
  const DifficultyContainer = difficultyContainers[difficulty];
  const difficultyLabel = difficultyLabels[difficulty];

  return (
    <Container>
      <Left>
        <ProblemLink>
          <a
            target="_blank"
            href={ `https://leetcode.com/problems/${slug}/` }
          >
            { `${id}. ${title}` }
          </a>
        </ProblemLink>

        <div className="ui horizontal divider">
          <DifficultyContainer
            style={ {
              color        : '#fff',
              fontSize     : '0.9rem',
              padding      : '0 2px',
              borderRadius : 3,
              marginLeft   : 8
            } }
          >
            { difficultyLabel }
          </DifficultyContainer>
        </div>

        <div className="ui stacked segment">
          <a className="ui blue left ribbon label">题意</a>
          <div dangerouslySetInnerHTML={ { __html: marked(explanation) } } />
        </div>

        <div className="ui piled segment">
          <a className="ui teal left ribbon label">解答</a>
          <div dangerouslySetInnerHTML={ { __html: marked(algorithm) } } />
        </div>
      </Left>

      <Right>
        <pre><code lang="golang">{ solution }</code></pre>
      </Right>
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
