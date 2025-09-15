# 状態管理: Jotai

Oreore.meでは、Jotaiを使用したアトミック状態管理により、効率的で保守性の高い状態管理を実現しています。

## 🌊 Jotaiの特徴

### Why Jotai?

| 特徴 | Jotai | Redux | Zustand |
|------|-------|--------|---------|
| 学習コスト | 低 | 高 | 中 |
| ボイラープレート | 最小 | 多 | 少 |
| TypeScript対応 | 優秀 | 良好 | 良好 |
| DevTools | 有り | 充実 | 有り |
| React統合 | 密結合 | 疎結合 | 疎結合 |
| パフォーマンス | 優秀 | 良好 | 良好 |

### アトミック設計の利点

```tsx
// ❌ Redux の場合 - 多くのボイラープレート
interface RootState {
  auth: AuthState;
  ui: UIState;
  // ...
}

const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    setUser: (state, action) => {
      state.user = action.payload;
    },
    setLoading: (state, action) => {
      state.loading = action.payload;
    },
  },
});

// ✅ Jotai の場合 - シンプルな定義
export const userAtom = atom<User | null>(null);
export const isLoadingAtom = atom(false);
export const isLoggedInAtom = atom((get) => get(userAtom) !== null);
```

## 🏗️ 状態構造設計

### Atom組織化

```
utils/atoms/
├── auth.ts              # 認証関連状態
├── ui.ts                # UI状態（モーダル、トースト等）
├── user.ts              # ユーザー情報
├── clients.ts           # OIDCクライアント状態
├── settings.ts          # アプリケーション設定
└── index.ts             # エクスポート統合
```

### 認証状態管理

```tsx
// utils/atoms/auth.ts
import { atom } from 'jotai';
import { atomWithStorage } from 'jotai/utils';

// 基本的な認証状態
export const currentUserAtom = atom<User | null>(null);

export const isLoggedInAtom = atom((get) => {
  const user = get(currentUserAtom);
  return user !== null;
});

export const loginStateAtom = atom<'loading' | 'authenticated' | 'unauthenticated'>((get) => {
  const user = get(currentUserAtom);
  if (user === undefined) return 'loading'; // 初期化中
  if (user === null) return 'unauthenticated';
  return 'authenticated';
});

// 永続化が必要な状態
export const rememberMeAtom = atomWithStorage('remember-me', false);

// 派生状態（計算プロパティ）
export const userPermissionsAtom = atom((get) => {
  const user = get(currentUserAtom);
  if (!user) return [];

  const permissions = ['read'];
  if (user.role === 'admin') permissions.push('admin');
  if (user.role === 'staff') permissions.push('manage');

  return permissions;
});

// 非同期アトム（ユーザー情報取得）
export const fetchCurrentUserAtom = atom(
  null,
  async (get, set) => {
    try {
      set(userLoadingAtom, true);
      const response = await fetch('/api/v2/user/me', {
        credentials: 'include',
      });

      if (response.ok) {
        const user = await response.json();
        set(currentUserAtom, user);
      } else {
        set(currentUserAtom, null);
      }
    } catch (error) {
      console.error('Failed to fetch user:', error);
      set(currentUserAtom, null);
    } finally {
      set(userLoadingAtom, false);
    }
  }
);

export const userLoadingAtom = atom(false);
```

### UI状態管理

```tsx
// utils/atoms/ui.ts
import { atom } from 'jotai';

// モーダル状態
export const activeModalAtom = atom<string | null>(null);

export const openModalAtom = atom(
  null,
  (get, set, modalId: string) => {
    set(activeModalAtom, modalId);
  }
);

export const closeModalAtom = atom(
  null,
  (get, set) => {
    set(activeModalAtom, null);
  }
);

// トースト通知
interface ToastMessage {
  id: string;
  title: string;
  description?: string;
  status: 'success' | 'error' | 'warning' | 'info';
  duration?: number;
}

export const toastMessagesAtom = atom<ToastMessage[]>([]);

export const addToastAtom = atom(
  null,
  (get, set, toast: Omit<ToastMessage, 'id'>) => {
    const id = Date.now().toString();
    const newToast: ToastMessage = { ...toast, id };
    const currentToasts = get(toastMessagesAtom);
    set(toastMessagesAtom, [...currentToasts, newToast]);

    // 自動削除
    const duration = toast.duration || 5000;
    setTimeout(() => {
      const toasts = get(toastMessagesAtom);
      set(toastMessagesAtom, toasts.filter(t => t.id !== id));
    }, duration);
  }
);

// サイドバー状態
export const sidebarOpenAtom = atomWithStorage('sidebar-open', true);

// テーマ設定
export const colorModeAtom = atomWithStorage('color-mode', 'light');
```

### 複雑な状態管理

```tsx
// utils/atoms/clients.ts
import { atom } from 'jotai';
import { focusAtom } from 'jotai/optics';

// OIDCクライアント管理
export const clientsAtom = atom<OIDCClient[]>([]);

export const selectedClientIdAtom = atom<string | null>(null);

// 選択中のクライアント（派生状態）
export const selectedClientAtom = atom((get) => {
  const clients = get(clientsAtom);
  const selectedId = get(selectedClientIdAtom);
  return clients.find(client => client.id === selectedId) || null;
});

// フィルタリングとソート
export const clientFilterAtom = atom('');
export const clientSortAtom = atom<'name' | 'created' | 'updated'>('created');

export const filteredClientsAtom = atom((get) => {
  const clients = get(clientsAtom);
  const filter = get(clientFilterAtom);
  const sort = get(clientSortAtom);

  let filtered = clients;

  // フィルタリング
  if (filter) {
    filtered = filtered.filter(client =>
      client.name.toLowerCase().includes(filter.toLowerCase()) ||
      client.description?.toLowerCase().includes(filter.toLowerCase())
    );
  }

  // ソート
  return filtered.sort((a, b) => {
    switch (sort) {
      case 'name':
        return a.name.localeCompare(b.name);
      case 'created':
        return new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime();
      case 'updated':
        return new Date(b.updatedAt).getTime() - new Date(a.updatedAt).getTime();
      default:
        return 0;
    }
  });
});

// クライアント作成/更新
export const createClientAtom = atom(
  null,
  async (get, set, clientData: CreateClientRequest) => {
    try {
      const response = await fetch('/api/v2/clients', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(clientData),
        credentials: 'include',
      });

      if (response.ok) {
        const newClient = await response.json();
        const currentClients = get(clientsAtom);
        set(clientsAtom, [...currentClients, newClient]);
        return newClient;
      } else {
        throw new Error('Failed to create client');
      }
    } catch (error) {
      console.error('Create client error:', error);
      throw error;
    }
  }
);

// 個別クライアントプロパティのFocusAtom
export const selectedClientNameAtom = focusAtom(selectedClientAtom, (optic) =>
  optic.prop('name')
);
```

## 🎯 コンポーネント統合

### Hookパターンの活用

```tsx
// hooks/useAuth.ts
import { useAtom, useAtomValue, useSetAtom } from 'jotai';
import {
  currentUserAtom,
  isLoggedInAtom,
  fetchCurrentUserAtom,
  loginStateAtom
} from '../atoms/auth';

export const useAuth = () => {
  const [currentUser, setCurrentUser] = useAtom(currentUserAtom);
  const isLoggedIn = useAtomValue(isLoggedInAtom);
  const loginState = useAtomValue(loginStateAtom);
  const fetchCurrentUser = useSetAtom(fetchCurrentUserAtom);

  const login = async (email: string, password: string) => {
    try {
      const response = await fetch('/api/v2/login/password', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
        credentials: 'include',
      });

      if (response.ok) {
        const user = await response.json();
        setCurrentUser(user);
        return { success: true, user };
      } else {
        const error = await response.json();
        return { success: false, error: error.message };
      }
    } catch (error) {
      return { success: false, error: 'Network error' };
    }
  };

  const logout = async () => {
    try {
      await fetch('/api/v2/logout', {
        method: 'POST',
        credentials: 'include',
      });
    } finally {
      setCurrentUser(null);
    }
  };

  return {
    currentUser,
    isLoggedIn,
    loginState,
    login,
    logout,
    fetchCurrentUser,
  };
};
```

### コンポーネントでの使用例

```tsx
// components/Auth/LoginForm.tsx
import { useAuth } from '../../hooks/useAuth';
import { useToast } from '../../hooks/useToast';

export const LoginForm: React.FC = () => {
  const { login } = useAuth();
  const { addToast } = useToast();
  const [isLoading, setIsLoading] = useState(false);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormData>();

  const onSubmit = async (data: LoginFormData) => {
    setIsLoading(true);

    try {
      const result = await login(data.email, data.password);

      if (result.success) {
        addToast({
          title: 'ログインしました',
          status: 'success',
        });
        // ページリダイレクト等
      } else {
        addToast({
          title: 'ログインに失敗しました',
          description: result.error,
          status: 'error',
        });
      }
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Box as="form" onSubmit={handleSubmit(onSubmit)}>
      <VStack spacing={4}>
        <FormControl isInvalid={!!errors.email}>
          <FormLabel>メールアドレス</FormLabel>
          <Input
            type="email"
            {...register('email', { required: 'メールアドレスは必須です' })}
          />
          <FormErrorMessage>{errors.email?.message}</FormErrorMessage>
        </FormControl>

        <FormControl isInvalid={!!errors.password}>
          <FormLabel>パスワード</FormLabel>
          <Input
            type="password"
            {...register('password', { required: 'パスワードは必須です' })}
          />
          <FormErrorMessage>{errors.password?.message}</FormErrorMessage>
        </FormControl>

        <Button
          type="submit"
          isLoading={isLoading}
          loadingText="ログイン中..."
          w="full"
        >
          ログイン
        </Button>
      </VStack>
    </Box>
  );
};
```

### プロバイダー設定

```tsx
// app/providers.tsx
'use client';

import { Provider as JotaiProvider, createStore } from 'jotai';
import { DevTools } from 'jotai-devtools';

// Jotaiストア作成
const store = createStore();

export function Providers({ children }: { children: React.ReactNode }) {
  return (
    <JotaiProvider store={store}>
      {/* 開発環境のみDevToolsを表示 */}
      {process.env.NODE_ENV === 'development' && <DevTools />}
      {children}
    </JotaiProvider>
  );
}
```

## 📡 SWR統合

### Jotai + SWR の組み合わせ

```tsx
// utils/atoms/swr.ts
import { atom } from 'jotai';
import useSWR from 'swr';

// SWR設定atom
export const swrConfigAtom = atom({
  revalidateOnFocus: false,
  revalidateOnReconnect: true,
  dedupingInterval: 60000,
});

// APIクライアント設定
export const apiClientAtom = atom({
  get: async (url: string) => {
    const response = await fetch(url, {
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
    });

    if (!response.ok) {
      throw new Error(`API Error: ${response.status}`);
    }

    return await response.json();
  },
});

// SWRをatomで活用するカスタムフック
export const useAtomSWR = <T>(key: string | null) => {
  const apiClient = useAtomValue(apiClientAtom);
  const config = useAtomValue(swrConfigAtom);

  return useSWR<T>(key, apiClient.get, config);
};
```

### SWR + Atom実装例

```tsx
// hooks/useClients.ts
export const useClients = () => {
  const { data, error, mutate } = useAtomSWR<OIDCClient[]>('/api/v2/clients');
  const [clients, setClients] = useAtom(clientsAtom);

  // SWRデータをAtomに同期
  useEffect(() => {
    if (data) {
      setClients(data);
    }
  }, [data, setClients]);

  const createClient = useSetAtom(createClientAtom);

  const handleCreateClient = async (clientData: CreateClientRequest) => {
    const newClient = await createClient(clientData);

    // SWRキャッシュを更新
    mutate([...clients, newClient], false);

    return newClient;
  };

  return {
    clients,
    isLoading: !error && !data,
    error,
    createClient: handleCreateClient,
    mutate,
  };
};
```

## 🧪 テスト戦略

### Atomテスト

```tsx
// __tests__/atoms/auth.test.ts
import { renderHook, act } from '@testing-library/react';
import { useAtom, useAtomValue } from 'jotai';
import { currentUserAtom, isLoggedInAtom } from '../atoms/auth';
import { JotaiTestProvider } from '../test-utils';

describe('Auth Atoms', () => {
  test('isLoggedInAtom は currentUserAtom に基づく', () => {
    const { result } = renderHook(() => ({
      user: useAtom(currentUserAtom),
      isLoggedIn: useAtomValue(isLoggedInAtom),
    }), {
      wrapper: JotaiTestProvider,
    });

    // 初期状態: ログインしていない
    expect(result.current.isLoggedIn).toBe(false);

    // ユーザー設定後: ログイン状態
    act(() => {
      const [, setUser] = result.current.user;
      setUser(mockUser);
    });

    expect(result.current.isLoggedIn).toBe(true);
  });
});
```

### Storybookでの状態管理

```tsx
// stories/decorators/JotaiDecorator.tsx
import { Provider as JotaiProvider, createStore } from 'jotai';
import { currentUserAtom } from '../../utils/atoms/auth';

export const JotaiDecorator = (Story: any, context: any) => {
  const store = createStore();

  // ストーリーパラメータから初期状態設定
  const jotaiState = context.parameters?.jotai || {};

  if (jotaiState.currentUser) {
    store.set(currentUserAtom, jotaiState.currentUser);
  }

  return (
    <JotaiProvider store={store}>
      <Story />
    </JotaiProvider>
  );
};

// 使用例
export const LoggedInState: Story = {
  parameters: {
    jotai: {
      currentUser: {
        id: '1',
        email: 'user@example.com',
        userName: 'testuser',
      },
    },
  },
  decorators: [JotaiDecorator],
};
```

## 🚀 パフォーマンス最適化

### 適切なAtom分割

```tsx
// ❌ 悪い例 - 巨大なオブジェクト
const userStateAtom = atom({
  profile: null,
  settings: {},
  preferences: {},
  notifications: [],
  // ... 他多数
});

// ✅ 良い例 - 関心事ごとに分離
export const userProfileAtom = atom<UserProfile | null>(null);
export const userSettingsAtom = atom<UserSettings>({});
export const userPreferencesAtom = atom<UserPreferences>({});
export const userNotificationsAtom = atom<Notification[]>([]);

// 必要な場合のみ組み合わせ
export const userSummaryAtom = atom((get) => ({
  profile: get(userProfileAtom),
  hasNotifications: get(userNotificationsAtom).length > 0,
}));
```

### メモ化されたAtom

```tsx
// 重い計算のメモ化
export const expensiveComputationAtom = atom((get) => {
  const data = get(largeDataSetAtom);

  // 重い処理のメモ化
  return useMemo(() => {
    return complexCalculation(data);
  }, [data]);
});

// 条件付きAtom
export const conditionalDataAtom = atom((get) => {
  const shouldFetch = get(shouldFetchDataAtom);

  if (!shouldFetch) {
    return null;
  }

  return get(expensiveDataAtom);
});
```

## 📚 関連ドキュメント

- [フロントエンドアーキテクチャ](./architecture.md)
- [コンポーネント設計](./components.md)
- [テスト戦略](./testing.md)