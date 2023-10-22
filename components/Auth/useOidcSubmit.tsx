interface Returns {
  submit: () => Promise<void>;
  cancel: () => Promise<void>;
}

export const useOidcSubmit = (): Returns => {
  const submit = async () => {
    window.alert('submit');
  };

  const cancel = async () => {
    window.alert('cancel');
  };

  return {
    submit,
    cancel,
  };
};
