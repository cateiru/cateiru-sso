# コンポーネント設計と構造

Oreore.meのフロントエンドは、機能別に整理された再利用可能なコンポーネント設計を採用しています。

## 📁 コンポーネント構造

```
components/
├── Auth/                       # 認証関連
│   ├── OidcRequirePage.tsx    # OIDC認可画面
│   ├── Consent.tsx            # 同意確認
│   ├── useOidcRequire.tsx     # OIDCフック
│   └── useWebAuthn.tsx        # WebAuthn統合
│
├── Common/                     # 共通コンポーネント
│   ├── Header.tsx             # ヘッダー
│   ├── Footer.tsx             # フッター
│   ├── Loading.tsx            # ローディング表示
│   ├── ErrorBoundary.tsx      # エラーハンドリング
│   └── Logo.tsx               # ロゴ
│
├── Login/                      # ログイン機能
│   ├── Login.tsx              # ログインメイン画面
│   ├── UserIDEmailPage.tsx    # メールアドレス入力
│   ├── OtpPage.tsx            # OTP認証
│   ├── LoginSuccess.tsx       # ログイン成功
│   └── useWebAuthn.tsx        # WebAuthn処理
│
├── RegisterAccount/            # アカウント登録
│   ├── RegisterAccountPage.tsx # 登録メイン画面
│   ├── Steps.tsx              # ステップ表示
│   ├── EmailResend.tsx        # メール再送信
│   └── RegisterForm.tsx       # 登録フォーム
│
├── Profile/                    # プロフィール管理
│   ├── ProfilePage.tsx        # プロフィール表示
│   ├── EditProfile.tsx        # プロフィール編集
│   └── AvatarUpload.tsx       # アバター更新
│
└── Client/                     # OIDCクライアント管理
    ├── ClientList.tsx         # クライアント一覧
    ├── ClientDetail.tsx       # クライアント詳細
    └── CreateClient.tsx       # クライアント作成
```

## 🏗️ コンポーネント設計原則

### 1. Single Responsibility Principle

各コンポーネントは単一の責任を持つ：

```tsx
// ❌ 悪い例 - 複数の責任
const UserDashboard = () => {
  const [user, setUser] = useState();
  const [clients, setClients] = useState();
  const [notifications, setNotifications] = useState();

  // ユーザー情報取得
  useEffect(() => { /* ... */ }, []);
  // クライアント情報取得
  useEffect(() => { /* ... */ }, []);
  // 通知取得
  useEffect(() => { /* ... */ }, []);

  return (
    <div>
      <UserProfile user={user} />
      <ClientList clients={clients} />
      <NotificationPanel notifications={notifications} />
    </div>
  );
};

// ✅ 良い例 - 単一責任
const UserDashboard = () => {
  return (
    <Container>
      <UserProfile />
      <ClientList />
      <NotificationPanel />
    </Container>
  );
};

const UserProfile = () => {
  const { user, isLoading } = useCurrentUser();

  if (isLoading) return <LoadingSpinner />;

  return <ProfileCard user={user} />;
};
```

### 2. Props Interface Design

明確で型安全なPropsインターフェース：

```tsx
// components/Common/Button.tsx
interface BaseButtonProps {
  children: React.ReactNode;
  variant?: 'primary' | 'secondary' | 'danger';
  size?: 'sm' | 'md' | 'lg';
  isLoading?: boolean;
  isDisabled?: boolean;
}

interface ButtonProps extends BaseButtonProps {
  onClick?: () => void;
  type?: 'button' | 'submit' | 'reset';
}

interface LinkButtonProps extends BaseButtonProps {
  href: string;
  external?: boolean;
}

export const Button: React.FC<ButtonProps> = ({
  children,
  variant = 'primary',
  size = 'md',
  isLoading = false,
  isDisabled = false,
  onClick,
  type = 'button',
}) => {
  return (
    <ChakraButton
      colorScheme={getColorScheme(variant)}
      size={size}
      isLoading={isLoading}
      isDisabled={isDisabled}
      onClick={onClick}
      type={type}
    >
      {children}
    </ChakraButton>
  );
};
```

### 3. Composition Pattern

コンポーネント合成によるフレキシビリティ：

```tsx
// components/Common/Card.tsx
interface CardProps {
  children: React.ReactNode;
  variant?: 'default' | 'outlined' | 'elevated';
}

const Card: React.FC<CardProps> = ({ children, variant = 'default' }) => {
  return (
    <Box
      bg="white"
      borderRadius="lg"
      border={variant === 'outlined' ? '1px solid' : 'none'}
      borderColor="gray.200"
      shadow={variant === 'elevated' ? 'md' : 'sm'}
    >
      {children}
    </Box>
  );
};

const CardHeader: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  return (
    <Box p={6} borderBottom="1px solid" borderColor="gray.100">
      {children}
    </Box>
  );
};

const CardBody: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  return <Box p={6}>{children}</Box>;
};

const CardFooter: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  return (
    <Box p={6} borderTop="1px solid" borderColor="gray.100">
      {children}
    </Box>
  );
};

// 使用例
const ProfileCard = () => {
  return (
    <Card variant="elevated">
      <CardHeader>
        <Heading size="md">プロフィール</Heading>
      </CardHeader>
      <CardBody>
        <ProfileDetails />
      </CardBody>
      <CardFooter>
        <Button>編集</Button>
      </CardFooter>
    </Card>
  );
};
```

## 🔐 認証コンポーネント

### WebAuthn統合

```tsx
// components/Auth/WebAuthnButton.tsx
interface WebAuthnButtonProps {
  mode: 'register' | 'login';
  onSuccess?: (result: AuthResult) => void;
  onError?: (error: Error) => void;
}

export const WebAuthnButton: React.FC<WebAuthnButtonProps> = ({
  mode,
  onSuccess,
  onError,
}) => {
  const [isLoading, setIsLoading] = useState(false);
  const toast = useToast();

  const handleWebAuthn = async () => {
    setIsLoading(true);

    try {
      const result = mode === 'register'
        ? await registerWithWebAuthn()
        : await loginWithWebAuthn();

      onSuccess?.(result);

      toast({
        title: mode === 'register' ? '登録完了' : 'ログイン成功',
        status: 'success',
      });
    } catch (error) {
      const errorMessage = getWebAuthnErrorMessage(error);
      onError?.(error);

      toast({
        title: 'エラーが発生しました',
        description: errorMessage,
        status: 'error',
      });
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Button
      leftIcon={<Icon as={FiFingerprint} />}
      onClick={handleWebAuthn}
      isLoading={isLoading}
      loadingText={mode === 'register' ? '登録中...' : 'ログイン中...'}
      size="lg"
      w="full"
    >
      {mode === 'register' ? 'パスキーで登録' : 'パスキーでログイン'}
    </Button>
  );
};

// WebAuthnエラーメッセージの統一化
const getWebAuthnErrorMessage = (error: any): string => {
  if (error.name === 'NotSupportedError') {
    return 'お使いのブラウザはWebAuthnに対応していません。';
  }
  if (error.name === 'SecurityError') {
    return 'セキュリティエラーが発生しました。HTTPSでアクセスしてください。';
  }
  if (error.name === 'NotAllowedError') {
    return 'ユーザーによりキャンセルされました。';
  }
  return '認証中にエラーが発生しました。';
};
```

### OTP入力コンポーネント

```tsx
// components/Auth/OtpInput.tsx
interface OtpInputProps {
  length: number;
  value: string;
  onChange: (value: string) => void;
  onComplete?: (value: string) => void;
  isLoading?: boolean;
  error?: string;
}

export const OtpInput: React.FC<OtpInputProps> = ({
  length,
  value,
  onChange,
  onComplete,
  isLoading = false,
  error,
}) => {
  const inputRefs = useRef<(HTMLInputElement | null)[]>([]);

  const handleChange = (index: number, inputValue: string) => {
    // 数字のみ許可
    if (!/^\d*$/.test(inputValue)) return;

    const newValue = value.split('');
    newValue[index] = inputValue;
    const updatedValue = newValue.join('');

    onChange(updatedValue);

    // 次の入力欄にフォーカス
    if (inputValue && index < length - 1) {
      inputRefs.current[index + 1]?.focus();
    }

    // 全て入力完了時
    if (updatedValue.length === length) {
      onComplete?.(updatedValue);
    }
  };

  const handleKeyDown = (index: number, e: React.KeyboardEvent) => {
    if (e.key === 'Backspace' && !value[index] && index > 0) {
      // 前の入力欄にフォーカス
      inputRefs.current[index - 1]?.focus();
    }
  };

  return (
    <FormControl isInvalid={!!error}>
      <FormLabel>認証コード</FormLabel>
      <HStack spacing={2} justify="center">
        {Array.from({ length }, (_, index) => (
          <Input
            key={index}
            ref={(el) => (inputRefs.current[index] = el)}
            value={value[index] || ''}
            onChange={(e) => handleChange(index, e.target.value)}
            onKeyDown={(e) => handleKeyDown(index, e)}
            maxLength={1}
            textAlign="center"
            size="lg"
            width="50px"
            height="50px"
            fontSize="xl"
            isDisabled={isLoading}
          />
        ))}
      </HStack>
      <FormErrorMessage>{error}</FormErrorMessage>
    </FormControl>
  );
};
```

## 📝 フォームコンポーネント

### 共通フォーム部品

```tsx
// components/Common/FormField.tsx
interface FormFieldProps {
  name: string;
  label: string;
  type?: 'text' | 'email' | 'password';
  placeholder?: string;
  helperText?: string;
  isRequired?: boolean;
  register: UseFormRegister<any>;
  error?: FieldError;
}

export const FormField: React.FC<FormFieldProps> = ({
  name,
  label,
  type = 'text',
  placeholder,
  helperText,
  isRequired = false,
  register,
  error,
}) => {
  return (
    <FormControl isInvalid={!!error} isRequired={isRequired}>
      <FormLabel>{label}</FormLabel>
      <Input
        type={type}
        placeholder={placeholder}
        {...register(name)}
      />
      {helperText && (
        <FormHelperText>{helperText}</FormHelperText>
      )}
      <FormErrorMessage>{error?.message}</FormErrorMessage>
    </FormControl>
  );
};

// 使用例
const RegisterForm = () => {
  const { register, handleSubmit, formState: { errors } } = useForm();

  return (
    <VStack as="form" spacing={4} onSubmit={handleSubmit(onSubmit)}>
      <FormField
        name="email"
        label="メールアドレス"
        type="email"
        placeholder="example@domain.com"
        isRequired
        register={register}
        error={errors.email}
      />

      <FormField
        name="userName"
        label="ユーザー名"
        placeholder="yamada-taro"
        helperText="半角英数字とハイフンが使用できます"
        isRequired
        register={register}
        error={errors.userName}
      />
    </VStack>
  );
};
```

## 📱 レスポンシブコンポーネント

### モバイル対応ナビゲーション

```tsx
// components/Common/Navigation.tsx
export const Navigation: React.FC = () => {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const { user } = useCurrentUser();

  const navItems = [
    { label: 'ダッシュボード', href: '/' },
    { label: 'プロフィール', href: '/profile' },
    { label: 'クライアント', href: '/clients' },
    { label: '設定', href: '/settings' },
  ];

  return (
    <>
      {/* デスクトップナビゲーション */}
      <HStack
        as="nav"
        spacing={8}
        display={{ base: 'none', md: 'flex' }}
      >
        {navItems.map((item) => (
          <Link
            key={item.href}
            href={item.href}
            fontWeight="medium"
            _hover={{ textDecoration: 'underline' }}
          >
            {item.label}
          </Link>
        ))}
      </HStack>

      {/* モバイルハンバーガーメニュー */}
      <IconButton
        display={{ base: 'block', md: 'none' }}
        onClick={onOpen}
        variant="ghost"
        icon={<HamburgerIcon />}
        aria-label="メニューを開く"
      />

      {/* モバイルドロワー */}
      <Drawer isOpen={isOpen} placement="right" onClose={onClose}>
        <DrawerOverlay />
        <DrawerContent>
          <DrawerCloseButton />
          <DrawerHeader>
            <HStack>
              <Avatar size="sm" src={user?.avatarUrl} />
              <Text>{user?.userName}</Text>
            </HStack>
          </DrawerHeader>

          <DrawerBody>
            <VStack align="stretch" spacing={4}>
              {navItems.map((item) => (
                <Button
                  key={item.href}
                  variant="ghost"
                  justifyContent="flex-start"
                  as={NextLink}
                  href={item.href}
                  onClick={onClose}
                >
                  {item.label}
                </Button>
              ))}
            </VStack>
          </DrawerBody>
        </DrawerContent>
      </Drawer>
    </>
  );
};
```

## 🎨 テーマ対応コンポーネント

### ダークモード対応

```tsx
// components/Common/ColorModeToggle.tsx
export const ColorModeToggle: React.FC = () => {
  const { colorMode, toggleColorMode } = useColorMode();

  return (
    <IconButton
      aria-label={colorMode === 'light' ? 'ダークモードに切り替え' : 'ライトモードに切り替え'}
      icon={colorMode === 'light' ? <MoonIcon /> : <SunIcon />}
      onClick={toggleColorMode}
      variant="ghost"
    />
  );
};

// テーマ対応スタイル
const ThemeAwareCard: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const bgColor = useColorModeValue('white', 'gray.800');
  const borderColor = useColorModeValue('gray.200', 'gray.600');

  return (
    <Box
      bg={bgColor}
      borderWidth="1px"
      borderColor={borderColor}
      borderRadius="lg"
      p={6}
    >
      {children}
    </Box>
  );
};
```

## 🧪 テスト可能な設計

### Storybook統合

```tsx
// components/Auth/WebAuthnButton.stories.tsx
export default {
  title: 'Auth/WebAuthnButton',
  component: WebAuthnButton,
  parameters: {
    backgrounds: {
      default: 'light',
    },
  },
} as Meta<typeof WebAuthnButton>;

type Story = StoryObj<typeof WebAuthnButton>;

export const Register: Story = {
  args: {
    mode: 'register',
  },
};

export const Login: Story = {
  args: {
    mode: 'login',
  },
};

export const Loading: Story = {
  args: {
    mode: 'register',
  },
  parameters: {
    msw: {
      handlers: [
        rest.post('/api/v2/register/begin_webauthn', (req, res, ctx) => {
          return res(ctx.delay('infinite'));
        }),
      ],
    },
  },
};

export const Error: Story = {
  args: {
    mode: 'register',
  },
  parameters: {
    msw: {
      handlers: [
        rest.post('/api/v2/register/begin_webauthn', (req, res, ctx) => {
          return res(ctx.status(500), ctx.json({ error: 'Server error' }));
        }),
      ],
    },
  },
};
```

## 📊 パフォーマンス最適化

### メモ化とコード分割

```tsx
// 重いコンポーネントのメモ化
const ExpensiveComponent = React.memo<{ data: ComplexData }>(({ data }) => {
  const processedData = useMemo(() => {
    return expensiveCalculation(data);
  }, [data]);

  return <div>{/* レンダリング */}</div>;
});

// 条件付きコンポーネント読み込み
const LazyAdminPanel = lazy(() => import('./AdminPanel'));

const Dashboard: React.FC = () => {
  const { user } = useCurrentUser();

  return (
    <div>
      <h1>ダッシュボード</h1>
      {user?.isAdmin && (
        <Suspense fallback={<LoadingSpinner />}>
          <LazyAdminPanel />
        </Suspense>
      )}
    </div>
  );
};
```

## 📚 関連ドキュメント

- [フロントエンドアーキテクチャ](./architecture.md)
- [状態管理](./state-management.md)
- [テスト戦略](./testing.md)