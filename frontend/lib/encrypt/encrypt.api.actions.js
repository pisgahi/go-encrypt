import { useQuery, useQueryClient } from "@tanstack/react-query";
import axiosInstance from "../axios/axios-instance";

// GET SECRET
export const getSecret = async ({ key }) => {
  const response = await axiosInstance.post("/get", {
    key: key,
  });
  return response.data;
};

export const useGetData = () => {
  return useQuery(["getSecret"], getSecret);
};

// CREATE SECRET
export const createSecret = async ({ plainText, key }) => {
  const response = await axiosInstance.post("/add", {
    plainText: plainText,
    key: key,
  });
  return response.data;
};

export const useCreateSecret = () => {
  const queryClient = useQueryClient();

  return useMutation(createSecret, {
    onSuccess: () => {
      queryClient.invalidateQueries(["createSecret"]);
    },
  });
};
