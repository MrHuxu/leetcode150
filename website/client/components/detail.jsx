import React from 'react';
import { connect } from 'react-redux';
import { string } from 'prop-types';
import marked from 'marked';

const Detail = ({ solution, explanation }) => {
  return (
    <div>
      <pre><code lang="golang">{ solution }</code></pre>
      <div dangerouslySetInnerHTML={ { __html: marked(explanation) } } />
    </div>
  );
};

Detail.propTypes = {
  solution    : string.isRequired,
  explanation : string.isRequired
};

const mapStateToProps = state => state;

export default connect(mapStateToProps)(Detail);
