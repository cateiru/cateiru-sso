import React from 'react';
import {OrgDetail} from '../../../../components/Staff/Org/OrgDetail';
import {StaffFrame} from '../../../../components/Staff/StaffFrame';

interface Props {
  params: {
    id: string;
  };
}

const Page: React.FC<Props> = ({params}) => {
  return (
    <StaffFrame
      title="組織詳細"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {href: '/staff/orgs', pageName: '組織一覧'},
        {pageName: '組織詳細'},
      ]}
    >
      <OrgDetail id={params.id} />
    </StaffFrame>
  );
};

export default Page;
