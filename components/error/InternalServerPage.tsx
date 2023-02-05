import {Text, Flex} from '@chakra-ui/react';

const InternalServerPage = () => {
  return (
    <Flex
      justifyContent="center"
      alignItems="center"
      flexDirection="column"
      width="100%"
      height="80vh"
      px={{base: '1rem', md: '5rem'}}
    >
      <Text
        fontSize={{base: '1.8rem', sm: '2.5rem'}}
        fontWeight="light"
        marginBottom="1.5rem"
      >
        500 | Internal Server Error.
      </Text>
      <Text fontSize={{base: '1rem', sm: '1.5rem'}} fontWeight="light">
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Λ＿Λ&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;＼＼
        <br />
        &nbsp;&nbsp;&nbsp;&nbsp;（&nbsp;&nbsp;&nbsp;&nbsp;・∀・）&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;ｶﾞｯ
        <br />
        &nbsp;&nbsp;&nbsp;と&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;）&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|
        <br />
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Ｙ&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;/ノ&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;人
        <br />
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;/&nbsp;&nbsp;）&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&lt;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&gt;_Λ∩
        <br />
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;＿/し&gt;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;／／.
        Ｖ｀Д´）/
        <br />
        （＿フ彡
      </Text>
    </Flex>
  );
};

export default InternalServerPage;
