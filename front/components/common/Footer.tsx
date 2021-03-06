import {
  Box,
  Stack,
  Center,
  Heading,
  StackDivider,
  Link,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';
import Logo from './Logo';

const Footer = () => {
  const user = useRecoilValue(UserState);

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
            <Box width="10rem">
              <NextLink href="/" passHref>
                <Link _focus={{boxShadow: 'none'}}>
                  <Logo />
                </Link>
              </NextLink>
            </Box>
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
                  {
                    text: 'お問い合わせ',
                    links: user
                      ? `https://cateiru.com/contact?url=https://sso.cateiru.com&name=${user?.last_name}%20${user?.first_name}&mail=${user?.mail}`
                      : 'https://cateiru.com/contact?url=https://sso.cateiru.com',
                  },
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
          <NextLink href="https://cateiru.com" passHref>
            <Link>&copy; {new Date().getFullYear()} cateiru</Link>
          </NextLink>
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
