package rpc
import(
  "context"
  "go_projects/solana"
)

func (cl *Client) GetBalance(ctx context.Context, account solana.PublicKey, commitment CommitmentType) (out *GetBalanceResult, err error){
  params := []interface{}{account}
  if commitment != ""{
    params = append(params, map[string]interface{}{"commitment": commitment})
  }
  err = cl.rpcClient.CallForInfo(ctx, &out, "getBalance", params)
  return 
}

