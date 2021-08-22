import {LocalSigner, Signer} from '@marsoccupied/plasma-client/lib/crypto/Signer';

export function signerFromStr (key: Buffer): Signer {
  return new LocalSigner(key);
}