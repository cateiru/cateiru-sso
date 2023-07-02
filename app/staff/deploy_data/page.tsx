import React from 'react';
import {DeployData} from '../../../components/Staff/DeployData';
import {StaffFrame} from '../../../components/Staff/StaffFrame';

const Page = () => {
  return (
    <StaffFrame
      title="デプロイデータ"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {pageName: 'デプロイデータ'},
      ]}
    >
      <DeployData />
    </StaffFrame>
  );
};

export default Page;
