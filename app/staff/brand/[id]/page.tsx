import React from 'react';
import {BrandDetail} from '../../../../components/Staff/Brand/BrandDetail';
import {StaffFrame} from '../../../../components/Staff/StaffFrame';

interface Props {
  params: {
    id: string;
  };
}

const Page: React.FC<Props> = ({params}) => {
  return (
    <StaffFrame
      title="ブランド詳細"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {href: '/staff/brands', pageName: 'ブランド一覧'},
        {pageName: 'ブランド詳細'},
      ]}
    >
      <BrandDetail id={params.id} />
    </StaffFrame>
  );
};

export default Page;
