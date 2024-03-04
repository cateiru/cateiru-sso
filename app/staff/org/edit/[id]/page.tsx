import React from 'react';
import {EditOrg} from '../../../../../components/Staff/Org/EditOrg';
import {StaffFrame} from '../../../../../components/Staff/StaffFrame';

interface Props {
  params: {
    id: string;
  };
}

const Page: React.FC<Props> = ({params}) => {
  return (
    <StaffFrame
      title="組織編集"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {href: '/staff/orgs', pageName: '組織一覧'},
        {href: `/staff/org/${params.id}`, pageName: '組織詳細'},
        {pageName: '組織編集'},
      ]}
    >
      <EditOrg id={params.id} />
    </StaffFrame>
  );
};

export default Page;
