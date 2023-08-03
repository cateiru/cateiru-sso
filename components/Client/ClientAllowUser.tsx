'use client';

import {
  Button,
  FormControl,
  FormErrorMessage,
  FormHelperText,
  FormLabel,
  Heading,
  Input,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  Tab,
  TabList,
  TabPanel,
  TabPanels,
  Tabs,
  Text,
  useColorModeValue,
  useDisclosure,
} from '@chakra-ui/react';
import {useParams} from 'next/navigation';
import React from 'react';
import {useForm} from 'react-hook-form';
import {useSWRConfig} from 'swr';
import {domainRegex} from '../../utils/regex';
import {Margin} from '../Common/Margin';
import {useSecondaryColor} from '../Common/useColor';
import {useRequest} from '../Common/useRequest';
import {ClientAllowUserTable} from './ClientAllowUserTable';

interface ClientAllowUserFormData {
  emailDomain?: string;
  userNameOrEmail?: string;
}

export const ClientAllowUser = () => {
  const {id} = useParams();

  const textColor = useSecondaryColor();
  const textBoldColor = useColorModeValue('gray.100', 'gray.800');

  const {isOpen, onOpen, onClose} = useDisclosure();
  const {request} = useRequest('/v2/client/allow_user');
  const {mutate} = useSWRConfig();

  const [tabIndex, setTabIndex] = React.useState(0);

  const {
    register,
    handleSubmit,
    reset,
    formState: {errors, isSubmitting},
  } = useForm<ClientAllowUserFormData>();

  const onSubmit = async (data: ClientAllowUserFormData) => {
    if (typeof id !== 'string') return;

    const form = new FormData();
    form.append('client_id', id);

    if (data.emailDomain) {
      form.append('email_domain', data.emailDomain);
    } else if (data.userNameOrEmail) {
      form.append('user_name_or_email', data.userNameOrEmail);
    }

    const res = await request({
      method: 'POST',
      body: form,
      mode: 'cors',
      credentials: 'include',
    });

    if (res) {
      reset();
      onClose();

      // パージする
      mutate(
        key =>
          typeof key === 'string' &&
          key.startsWith(`/v2/client/allow_user?client_id=${id}`),
        undefined,
        {revalidate: true}
      );
    }
  };

  return (
    <Margin>
      <Heading textAlign="center" mt="3rem">
        許可ユーザーの編集
      </Heading>
      <Text textAlign="center" mt="1rem" color={textColor}>
        許可するユーザーを追加することでそのユーザー、そのドメインのメールアドレスを持つユーザーのみ認証を許可することができます。
      </Text>
      <Button w="100%" mt="1.5rem" colorScheme="cateiru" onClick={onOpen}>
        ルールを追加
      </Button>
      <Modal isOpen={isOpen} onClose={onClose} isCentered>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>ルールを追加</ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody mb=".5rem">
            <form onSubmit={handleSubmit(onSubmit)}>
              <Tabs
                colorScheme="cateiru"
                isFitted
                onChange={index => {
                  reset();
                  setTabIndex(index);
                }}
              >
                <TabList>
                  <Tab>メールドメイン</Tab>
                  <Tab>ユーザー</Tab>
                </TabList>
                <TabPanels>
                  <TabPanel p="0">
                    <FormControl isInvalid={!!errors.emailDomain} mt="1rem">
                      <FormLabel htmlFor="emailDomain">
                        メールドメイン
                      </FormLabel>
                      <FormHelperText color={textColor} mb=".5rem">
                        例えば、
                        <Text
                          as="span"
                          bgColor={textBoldColor}
                          px=".3rem"
                          borderRadius="5px"
                        >
                          example.test
                        </Text>
                        を設定すると
                        <Text
                          as="span"
                          bgColor={textBoldColor}
                          px=".2rem"
                          borderRadius="5px"
                        >
                          example@example.test
                        </Text>
                        といったドメインのメールアドレスが許可されます。
                      </FormHelperText>
                      <Input
                        id="emailDomain"
                        type="text"
                        {...register('emailDomain', {
                          pattern: {
                            value: domainRegex,
                            message: '正しいドメインを入力してください',
                          },
                          required:
                            tabIndex === 0
                              ? 'メールドメインを入力してください。'
                              : undefined,
                        })}
                      />
                      <FormErrorMessage>
                        {errors.emailDomain && errors.emailDomain.message}
                      </FormErrorMessage>
                    </FormControl>
                  </TabPanel>
                  <TabPanel p="0">
                    <FormControl isInvalid={!!errors.userNameOrEmail} mt="1rem">
                      <FormLabel htmlFor="description">
                        ユーザー名またはメールアドレス
                      </FormLabel>
                      <FormHelperText color={textColor} mb=".5rem">
                        ユーザーを指定するとそのユーザーのみ認証を許可します。
                      </FormHelperText>
                      <Input
                        id="userNameOrEmail"
                        type="text"
                        autoComplete="username email"
                        {...register('userNameOrEmail', {
                          required:
                            tabIndex === 1
                              ? 'ユーザー名またはメールアドレスを入力してください。'
                              : undefined,
                        })}
                      />
                      <FormErrorMessage>
                        {errors.userNameOrEmail &&
                          errors.userNameOrEmail.message}
                      </FormErrorMessage>
                    </FormControl>
                  </TabPanel>
                </TabPanels>
              </Tabs>

              <Button
                mt={4}
                colorScheme="cateiru"
                isLoading={isSubmitting}
                type="submit"
                w="100%"
              >
                ルールを追加
              </Button>
            </form>
          </ModalBody>
        </ModalContent>
      </Modal>
      <ClientAllowUserTable id={id} />
    </Margin>
  );
};
