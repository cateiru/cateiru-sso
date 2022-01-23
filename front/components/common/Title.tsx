import Head from 'next/head';
import React from 'react';

const Title: React.FC<{title: string}> = ({title}) => {
  return (
    <Head>
      <title>{title || 'CateiruSSO'}</title>
    </Head>
  );
};

export default Title;
