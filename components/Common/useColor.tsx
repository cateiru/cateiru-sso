import {useColorModeValue} from '@chakra-ui/react';

// 目立たないテキストなど
export const useSecondaryColor = () => {
  return useColorModeValue('#718096', '#A0AEC0');
};

// 削除などのボタン
export const useDeleteColor = () => {
  return useColorModeValue('#F56565', '#E53E3E');
};

// ボーダー
export const useBorderColor = () => {
  return useColorModeValue('gray.300', 'gray.600');
};

// ドロップシャドウ
export const useShadowColor = () => {
  return useColorModeValue('#242838', '#000');
};
