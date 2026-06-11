import { beforeEach, describe, expect, it, vi } from 'vitest'

const { get, post } = vi.hoisted(() => ({
  get: vi.fn(),
  post: vi.fn(),
}))

vi.mock('@/api/client', () => ({
  apiClient: {
    get,
    post,
  },
}))

import { performUpdate } from '@/api/admin/system'

describe('admin system api', () => {
  beforeEach(() => {
    get.mockReset()
    post.mockReset()
  })

  it('uses an extended timeout for in-place system updates', async () => {
    const response = {
      message: 'Update completed. Please restart the service.',
      need_restart: true,
    }
    post.mockResolvedValue({ data: response })

    const result = await performUpdate()

    expect(post).toHaveBeenCalledWith('/admin/system/update', undefined, {
      timeout: 10 * 60 * 1000,
    })
    expect(result).toEqual(response)
  })
})
