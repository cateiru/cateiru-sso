import React from 'react';
import {EmailTemplatePreview} from '../../../components/Staff/EmailTemplatePreview';
import {StaffFrame} from '../../../components/Staff/StaffFrame';

const Page = () => {
  return (
    <StaffFrame
      title="メールテンプレートプレビュー"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {pageName: 'メールテンプレートプレビュー'},
      ]}
    >
      <EmailTemplatePreview />
    </StaffFrame>
  );
};

export default Page;
