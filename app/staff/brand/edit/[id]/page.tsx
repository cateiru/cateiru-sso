import React from 'react';
import {StaffFrame} from '../../../../../components/Staff/StaffFrame';

interface Props {
  params: {
    id: string;
  };
}

const Page: React.FC<Props> = ({params}) => {
  return (
    <StaffFrame
      title="ブランド編集"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {href: '/staff/brands', pageName: 'ブランド一覧'},
        {href: `/staff/brand/${params.id}`, pageName: 'ブランド詳細'},
        {pageName: 'ブランド編集'},
      ]}
    >
      <></>
    </StaffFrame>
  );
};

export default Page;
