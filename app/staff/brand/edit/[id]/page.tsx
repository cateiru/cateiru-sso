import React from 'react';
import {EditBrand} from '../../../../../components/Staff/Brand/EditBrand';
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
      <EditBrand id={params.id} />
    </StaffFrame>
  );
};

export default Page;
