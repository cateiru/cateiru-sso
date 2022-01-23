import {
  Box,
  Stack,
  Center,
  Text,
  Heading,
  StackDivider,
  Link,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import React from 'react';

const Footer = () => {
  return (
    <Box
      as="footer"
      role="contentinfo"
      mx="auto"
      py="12"
      px={{base: '4', md: '8'}}
    >
      <Stack spacing="10" divider={<StackDivider />}>
        <Stack
          direction={{base: 'column', lg: 'row'}}
          spacing={{base: '10', lg: '28'}}
        >
          <Box flex="1">
            <Link to="/" _focus={{boxShadow: 'none'}}>
              <Box width="10rem">TODO: LOGO</Box>
            </Link>
          </Box>
          <Stack
            direction={{
              base: 'column',
              sm: 'row',
              md: 'row',
              lg: 'row',
            }}
            spacing={{base: '10', lg: '28'}}
          >
            <Box>
              <FooterList
                title="About"
                elements={[
                  {text: 'CateiruSSOについて', links: '/about'},
                  {text: 'よくある質問', links: '/question'},
                  {text: '変更履歴', links: '/changelog'},
                  {text: '使い方', links: '/usage'},
                ]}
              />
            </Box>
            <Box>
              <FooterList
                title="Legal"
                elements={[
                  {text: '利用規約', links: '/terms'},
                  {text: 'プライバシーポリシー', links: '/privacy'},
                ]}
              />
            </Box>
            <Box>
              <FooterList
                title="Links"
                elements={[
                  {
                    text: 'GitHub',
                    links: 'https://github.com/cateiru/cateiru-sso',
                    isExternal: true,
                  },
                ]}
              />
            </Box>
          </Stack>
        </Stack>
        <Center>
          <Text fontSize="sm">&copy; {new Date().getFullYear()} cateiru</Text>
        </Center>
      </Stack>
    </Box>
  );
};

const FooterList = React.memo<{
  title: string;
  elements: {
    text: string;
    links: string;
    isExternal?: boolean;
  }[];
}>(({title, elements}) => {
  return (
    <Box>
      <Heading as="h4" fontSize="1.1rem" letterSpacing="wider" mb={4}>
        {title}
      </Heading>
      <Stack fontSize=".9rem">
        {elements.map((value, index) => {
          if (value.isExternal) {
            return (
              <Link
                href={value.links}
                key={index}
                isExternal={value.isExternal}
              >
                {value.text}
              </Link>
            );
          } else {
            return (
              <NextLink href={value.links} key={index} passHref>
                <Link>{value.text}</Link>
              </NextLink>
            );
          }
        })}
      </Stack>
    </Box>
  );
});

FooterList.displayName = 'footerName';

export default Footer;
